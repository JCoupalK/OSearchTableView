package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"os"
)

type Config struct {
	URL       string `json:"url"`
	User      string `json:"user"`
	Password  string `json:"password"`
	IndexName string `json:"index_name"`
	Size      int    `json:"size"`
	CSVFile   string `json:"csv_file"`
}

func main() {
	// Define flags with both long and short forms
	var (
		url           = flag.String("url", "", "OpenSearch URL (short form: -u)")
		urlShort      = flag.String("u", "", "OpenSearch URL (short form of -url)")
		user          = flag.String("user", "", "OpenSearch User (short form: -U)")
		userShort     = flag.String("U", "", "OpenSearch User (short form of -user)")
		password      = flag.String("password", "", "OpenSearch Password (short form: -p)")
		passwordShort = flag.String("p", "", "OpenSearch Password (short form of -password)")
		indexName     = flag.String("index", "", "Index Name (short form: -i)")
		indexShort    = flag.String("i", "", "Index Name (short form of -index)")
		size          = flag.Int("size", 0, "Size limit for the number of documents to fetch (short form: -s)")
		sizeShort     = flag.Int("s", 10, "Size limit for the number of documents to fetch (short form: -s)")
		configFile    = flag.String("config", "", "Config file path (short form: -c)")
		configShort   = flag.String("c", "", "Config file path (short form of -config)")
		csvFile       = flag.String("csv", "", "CSV output file (short form: -o)")
		csvShort      = flag.String("o", "", "CSV output file (short form of -csv)")
	)
	// Override the default flag.Usage
	flag.Usage = Usage

	flag.Parse()

	// Use the short form values if the long form was not provided
	if *url == "" {
		url = urlShort
	}
	if *user == "" {
		user = userShort
	}
	if *password == "" {
		password = passwordShort
	}
	if *indexName == "" {
		indexName = indexShort
	}
	if *size == 0 {
		size = sizeShort
	}
	if *configFile == "" {
		configFile = configShort
	}
	if *csvFile == "" {
		csvFile = csvShort
	}

	// Load config from file if provided
	var config Config
	if *configFile != "" {
		file, err := os.ReadFile(*configFile)
		if err != nil {
			fmt.Println("Error reading config file:", err)
			return
		}
		err = json.Unmarshal(file, &config)
		if err != nil {
			fmt.Println("Error parsing config file:", err)
			return
		}
	} else {
		config = Config{
			URL:       *url,
			User:      *user,
			Password:  *password,
			IndexName: *indexName,
			Size:      *size,
			CSVFile:   *csvFile,
		}
	}

	// Validate required config
	if config.URL == "" || config.IndexName == "" {
		fmt.Println("\nURL and Index Name are required.")
		return
	}

	if config.Size > 10000 {
		fmt.Println("\nMaximum size is 10000 by OpenSearch API limitations.")
		return
	}

	// Fetch data from OpenSearch
	FetchData(config)
}
