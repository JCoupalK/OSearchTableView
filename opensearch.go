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
		"sort": []map[string]interface{}{
			{
				config.TimestampField: map[string]interface{}{
					"order":         "desc",
					"unmapped_type": "date",
				},
			},
		},
	}
	b, _ := json.Marshal(query)

	// DEBUG:
	// fmt.Printf("Generated Query: %s\n", string(b))

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

	// DEBUG:
	// if res.StatusCode != 200 {
	// 	fmt.Printf("Unexpected status code: %d\n", res.StatusCode)
	// 	bodyBytes, _ := io.ReadAll(res.Body)
	// 	fmt.Printf("Response: %s\n", string(bodyBytes))
	// 	return
	// }

	var r map[string]interface{}
	if err := json.NewDecoder(res.Body).Decode(&r); err != nil {
		fmt.Printf("Error parsing the response body: %s\n", err)
		return
	}

	// Check if "hits" exist in the response
	hitsRaw, ok := r["hits"]
	if !ok {
		fmt.Println("No hits found in the response.")
		return
	}

	hitsMap, ok := hitsRaw.(map[string]interface{})
	if !ok {
		fmt.Println("Error: Unexpected format for hits.")
		return
	}

	hits, ok := hitsMap["hits"].([]interface{})
	if !ok || len(hits) == 0 {
		fmt.Println("No data found.")
		return
	}

	firstHit, ok := hits[0].(map[string]interface{})
	if !ok {
		fmt.Println("Error: Unexpected format for hit.")
		return
	}

	source, ok := firstHit["_source"].(map[string]interface{})
	if !ok {
		fmt.Println("Error: Unexpected format for _source.")
		return
	}

	headers := make([]string, 0, len(source))
	for field := range source {
		headers = append(headers, field)
	}

	var wg sync.WaitGroup
	rowsChan := make(chan []string, len(hits))

	for _, hit := range hits {
		wg.Add(1)
		go func(hit interface{}) {
			defer wg.Done()
			row := make([]string, len(headers))
			doc, ok := hit.(map[string]interface{})["_source"].(map[string]interface{})
			if !ok {
				fmt.Println("Error: Unexpected format for document source.")
				return
			}
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
