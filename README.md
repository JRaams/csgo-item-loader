# CS:GO Item Loader

## 1. Description

Parses CS:GO items from game files into api-ready json.

## 2. Requirements

Golang (v1.15.8): https://golang.org/dl/

## 3. Usage

#### 3.1 Building the app

```bash
# Creates the csgo-item-loader executable
$ go build 
``` 

#### 3.2 Commands 

##### 3.2.1 Downloading game files

```bash
Command: 'download' (alias 'd')
$ ./csgo-item-loader download
```

Result:
`./storage` folder with the following contents:
- `csgo_english.txt`
- `items_game_cdn.txt`
- `items_game.txt`

##### 3.2.2 Parsing game files into usable item data (todo)

- [x] Parsing collections and checking that they correspond with the collections in `resources/collections.json`
- [ ] Parsing weapons and weapontypes
- [ ] Parsing paintkits
- [ ] Parsing knives
- [ ] Parsing gloves

```bash
Command 'parse' (alias 'p')
$ ./csgo-item-loader parse
```

Result:
`./resources` folder with the following contents:
- [x] `collections.json` (Not automatically updated)
- [ ] `weapons.json`
- [ ] `paintkits.json`
- [ ] `knives.json`
- [ ] `gloves.json`

##### 3.2.3 Stitching game data together (todo)

Stitching weapons and paintkits together to create `resources/skins.json`

```bash
Command 'stitch' (alias 's')
$ ./csgo-item-loader stitch
```
