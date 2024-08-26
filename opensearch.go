package main

import (
	"bytes"
	"context"
	"crypto/tls"
	"encoding/csv"
	"encoding/json"
	"fmt"
	"net/http"
	"os"
	"sync"

	"github.com/olekukonko/tablewriter"
	"github.com/opensearch-project/opensearch-go"
)

func FetchData(config Config) {
	// Configure the OpenSearch client
	cfg := opensearch.Config{
		Transport: &http.Transport{
			TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
		},
		Addresses: []string{config.URL},
		Username:  config.User,
		Password:  config.Password,
	}

	client, err := opensearch.NewClient(cfg)
	if err != nil {
		fmt.Println("Error creating the client:", err)
		return
	}

	query := map[string]interface{}{
		"size": config.Size,
		"query": map[string]interface{}{
			"match_all": map[string]interface{}{},
		},
	}
	b, _ := json.Marshal(query)

	res, err := client.Search(
		client.Search.WithContext(context.Background()),
		client.Search.WithIndex(config.IndexName),
		client.Search.WithBody(bytes.NewReader(b)),
		client.Search.WithPretty(),
	)
	if err != nil {
		fmt.Println("Error getting response:", err)
		return
	}
	defer res.Body.Close()

	var r map[string]interface{}
	if err := json.NewDecoder(res.Body).Decode(&r); err != nil {
		fmt.Printf("Error parsing the response body: %s\n", err)
		return
	}

	hits := r["hits"].(map[string]interface{})["hits"].([]interface{})
	if len(hits) == 0 {
		fmt.Println("No data found.")
		return
	}

	firstHitSource := hits[0].(map[string]interface{})["_source"].(map[string]interface{})
	headers := make([]string, 0, len(firstHitSource))
	for field := range firstHitSource {
		headers = append(headers, field)
	}

	var wg sync.WaitGroup
	rowsChan := make(chan []string, len(hits))

	for _, hit := range hits {
		wg.Add(1)
		go func(hit interface{}) {
			defer wg.Done()
			row := make([]string, len(headers))
			doc := hit.(map[string]interface{})["_source"].(map[string]interface{})
			for i, header := range headers {
				if value, ok := doc[header]; ok {
					row[i] = fmt.Sprintf("%v", value)
				} else {
					row[i] = "N/A"
				}
			}
			rowsChan <- row
		}(hit)
	}

	go func() {
		wg.Wait()
		close(rowsChan)
	}()

	if config.CSVFile != "" {
		file, err := os.Create(config.CSVFile)
		if err != nil {
			fmt.Printf("Error creating CSV file: %s\n", err)
			return
		}
		defer file.Close()

		writer := csv.NewWriter(file)
		defer writer.Flush()

		// Write headers to the CSV file
		writer.Write(headers)

		// Write rows to the CSV file
		for row := range rowsChan {
			writer.Write(row)
		}

		fmt.Printf("Data successfully written to %s\n", config.CSVFile)
	} else {
		table := tablewriter.NewWriter(os.Stdout)
		table.SetHeader(headers)

		for row := range rowsChan {
			table.Append(row)
		}

		table.Render()
	}
}
