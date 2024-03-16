# OSearchTableView

OSearchTableView is a command-line tool written in Go for fetching and displaying data from OpenSearch indices in a tabular format.

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
- Support for pagination to fetch large datasets efficiently.
- Configurable options including URL, user authentication, index name, and more.

## Installation

1. Download the binary with wget:

    ```shell
    wget https://github.com/JCoupalK/OSearchTableView/releases/download/1.0/ostableview_linux_amd64_1.0.tar.gz
    ```

2. Unpack it with tar

    ```shell
    tar -xf ostableview_linux_amd64_1.0.tar.gz
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

```bash
Usage: ostableview [options]
    General options:
      -u,    --url              OpenSearch URL
      -U,    --user              OpenSearch user
      -p,    --password         OpenSearch password
      -i,    --index              Index name
      -s,    --size          Size limit for the number of rows to fetch (Default is 10, Maximum is 10000)
      -c,    --config          Config file path (replaces above arguments)
```

## Examples

```bash
# locally hosted opensearch with 100 rows queried.
./ostableview -u <http://localhost:9200> -U demo-user -p demo-password -i demo_index -s 100

# print output to file called results.txt
./ostableview -u <http://localhost:9200> -U demo-user -p demo-password -i demo_index -s 100 > results.txt

# configuration file specified
./ostableview -c path/to/config.json
```

Example of config.json:

```json
{
  "url": "http://localhost:9200",
  "user": "demo-user",
  "password": "demo-passwor",
  "index_name": "demo_index",
  "size": 100
}
```

## Contributing

Contributions are welcome. If you find a bug or have a feature request, please open an issue or submit a pull request.

## License

This project is licensed under the MIT License - see the LICENSE file for details.
