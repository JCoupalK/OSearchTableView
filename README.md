# OSearchTableView

OSearchTableView is a command-line tool for fetching and displaying data from OpenSearch indices in a tabular text format or export to a CSV file.

## Table of Contents

- [Features](#features)
- [Installation](#installation)
- [Building from Source](#building-from-source)
- [Usage](#usage)
- [Examples](#examples)
- [Contributing](#contributing)
- [License](#license)

## Features

- Fetch data from OpenSearch indices.
- Display fetched data in rows and tables.
- Export fetched data to a CSV file.
- Query the latest data using a customizable timestamp field.
- Uses concurrency through goroutines to fetch large datasets efficiently.
- Configurable options including URL, user authentication, index name, and more.

## Installation

1. Download the Linux amd64 binary with wget (more versions on [release](https://github.com/JCoupalK/OSearchTableView/releases/tag/1.1) tab):

    ```shell
    wget https://github.com/JCoupalK/OSearchTableView/releases/download/1.0/ostableview_linux_amd64_1.1.tar.gz
    ```

2. Unpack it with tar

    ```shell
    tar -xf ostableview_linux_amd64_1.1.tar.gz
    ```

3. Move it to your /usr/local/bin/ (Optional):

    ```shell
    sudo mv ostableview /usr/local/bin/ostableview
    ```

## Building from Source

1. Ensure you have Go installed on your system. You can download Go from [here](https://go.dev/dl/).
2. Clone the repository:

    ```shell
    git clone https://github.com/JCoupalK/OSearchTableView
    ```

3. Navigate to the cloned directory:

    ```shell
    cd OSearchTableView
    ```

4. Build the tool:

    ```shell
    go build -o ostableview .
    ```

## Usage

```text
Usage: ostableview [options]
    General options:
      -u,    --url                  OpenSearch URL
      -U,    --user                 OpenSearch user
      -p,    --password             OpenSearch password
      -i,    --index                Index name
      -s,    --size                 Size limit for the number of rows to fetch (Default is 10, Maximum is 10000)
      -o,    --csv                  CSV output file (if specified, data will be written to this file instead of displayed in a table)
      -t,    --timestamp            Timestamp field name (default is @timestamp, used to sort the documents by latest first)
      -c,    --config               Config file path (replaces above arguments)
```

## Examples

```bash
# Locally hosted OpenSearch with 100 rows queried.
./ostableview -u http://localhost:9200 -U demo-user -p demo-password -i demo_index -s 100

# Print tabular output to file called results.txt
./ostableview -u http://localhost:9200 -U demo-user -p demo-password -i demo_index -s 100 > results.txt

# Export to a CSV file called results.csv
./ostableview -u http://localhost:9200 -U demo-user -p demo-password -i demo_index -s 100 -o results.csv

# Query using a custom timestamp field and export to CSV
./ostableview -u http://localhost:9200 -U demo-user -p demo-password -i demo_index -s 100 -t "Time" -o results.csv

# Configuration file specified
./ostableview -c path/to/config.json

# Configuration file without specified index and size inside the file
./ostableview -c path/to/config.json -i demo_index -s 100
```

Example of JSON config file (you can include only what you need and use CLI arguments for the rest):

```json
{
  "url": "http://localhost:9200",
  "user": "demo-user",
  "password": "demo-password",
  "index_name": "demo_index",
  "size": 100,
  "csv_file": "results.csv",
  "timestamp_field": "Time"
}
```

## Contributing

Contributions are welcome. If you find a bug or have a feature request, please open an issue or submit a pull request.

## License

This project is licensed under the MIT License - see the LICENSE file for details.
