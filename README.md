# CS:GO Item Loader

## Description

Parses CS:GO items from game files into api-ready json.

## Requirements

Golang (v1.15.8): https://golang.org/dl/

## Usage

```bash
# Build the app:
$ go build # -> csgo-item-loader executable

# Commands

# 1. Download game files to disk
- 'download' (alias 'd')
$ ./csgo-item-loader download

# 2. Parse game files into usable text files
- 'parse' (alias 'p')
$ ./csgo-item-loader parse

# 3. Stitch parsed data into json
- 'stitch' (alias 's')
$ ./csgo-item-loader stitch
```