# CS:GO Item Loader

[![Go Report Card](https://goreportcard.com/badge/github.com/JRaams/csgo-item-loader)](https://goreportcard.com/report/github.com/JRaams/csgo-item-loader) [![License](https://img.shields.io/badge/license-MIT-brightgreen)](./LICENSE)

Parses CS:GO weapons/skins from game files into api-ready json.

## Requirements

Golang (v1.15.8): https://golang.org/dl/

## Usage

```bash
$ go run .
NAME:
   csgo-item-loader - A new cli application

USAGE:
   csgo-item-loader [global options] command [command options] [arguments...]

DESCRIPTION:
   Parses CS:GO items from game files into api-ready json.

COMMANDS:
   download, d
   parse, p
   stitch, s
   help, h      Shows a list of commands or help for one command

GLOBAL OPTIONS:
   --help, -h  show help (default: false)
```

1. Download game files
   `go run . d`
   Writes to 'storage/'

2. Parse VDF into json
   `go run . p`
   Writes to 'resources/'

3. Stitch data
   `go run . s`
   Writes to 'output/'
