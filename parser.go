package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"strings"
)

func parse() {
	fmt.Println("Parsing collections files...")
	_ = parseCollections()
	fmt.Println("resources/collections.json is up to date! Parsing weapons...")

	// fmt.Println("Parsing weapons")
	// weapons := parseWeapons()
	// fmt.Println(weapons)
}

// Collection type
type Collection struct {
	ID       int    `json:"id"`
	Name     string `json:"name"`
	Released string `json:"released"`
	Stattrak bool   `json:"stattrak"`
	Souvenir bool   `json:"souvenir"`
	Tag      string `json:"tag"`
}

// Weapon type
type Weapon struct {
	ID   int
	Name string
	Type string
	Desc string
	Tag  string
}

func parseCollections() map[string]*Collection {
	// Load collections from resources
	collectionsBytes, err := ioutil.ReadFile("resources/collections.json")
	if err != nil {
		log.Fatalf("Error loading collections: %s", err.Error())
	}
	collections := []*Collection{}
	json.Unmarshal(collectionsBytes, &collections)
	collectionsByName := map[string]*Collection{}
	for _, collection := range collections {
		collectionsByName[strings.ToLower(collection.Name)] = collection
	}

	// Check if all actual collections have an entry in resources/collections
	// If not, it needs to be added manually
	text := getTextBetweenLines("storage/csgo_english.txt", "// SET DESCRIPTIONS", "/////////////////")
	vdf := parseVdf(text)

	for key, value := range vdf {
		// Do not check descriptions or short names
		if strings.Contains(key, "_desc") || strings.Contains(key, "_short") {
			continue
		}
		if collectionsByName[strings.ToLower(key)] == nil {
			log.Fatalf("Error parsing collections: %s (%s) was not found in resources/collections.json, make sure to add it manually.", key, value)
		}
	}

	return collectionsByName
}

func parseWeapons() []*Weapon {
	text := getTextBetweenLines("storage/items_game.txt", "\"items\"", "\"attributes\"")
	fmt.Println(text)

	return nil
}

func parseVdf(lines []string) map[string]string {
	vdf := map[string]string{}

	for _, line := range lines {
		// Ignore empty lines or comments
		if len(line) < 2 || line[0:2] == "//" {
			continue
		}
		lineWithoutSpaces := strings.ReplaceAll(strings.ReplaceAll(line, " ", ""), "\t", "");

		parts := strings.Split(lineWithoutSpaces, `""`)
		if len(parts) == 1 {
			continue;
		}

		key := parts[0][1:]
		value := ""

		if len(strings.TrimSpace(key)) == 0 {
			continue
		}

		if len(parts) > 1 {
			value = parts[1][:len(parts[1])-1]
		}
		vdf[key] = value
	}

	return vdf
}

func getTextBetweenLines(file string, start string, end string) []string {
	f, err := os.Open(file)
	if err != nil {
		log.Fatal(err)
	}

	var lines []string
	startIndent := -1

	r := bufio.NewReader(f)
	bytes, _, err := r.ReadLine()
	for err == nil {
		line := string(bytes[:])
		if strings.TrimSpace(line) == start {
			startIndent = countLeadingSpace(string(bytes))
		}
		if startIndent != -1 {
			if strings.TrimSpace(line) == end {
				if endIndent := countLeadingSpace(string(bytes)); endIndent == startIndent {
					return lines
				}
			}
			lines = append(lines, line)
		}
		bytes, _, err = r.ReadLine()
	}

	return lines
}

func countLeadingSpace(line string) int {
	i := 0
	for _, rune := range line {
		if rune == '	' {
			i += 4
		} else if rune == ' ' {
			i++
		} else {
			break
		}
	}
	return i
}
