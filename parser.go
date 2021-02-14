package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"regexp"
	"strings"
)

func parse() {
	fmt.Println("Parsing game files!")
	parseCollections()
}

type Collection struct {
	ID       int    `json:"id"`
	Name     string `json:"name"`
	Released string `json:"released"`
	Stattrak bool   `json:"stattrak"`
	Souvenir bool   `json:"souvenir"`
	Tag      string `json:"tag"`
}

func parseCollections() {
	// Get actual collections from gamefiles
	text := getTextBetweenLines("storage/csgo_english.txt", "// SET DESCRIPTIONS", "/////////////////")
	vdf := parseVdf(text)

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
	for key, value := range vdf {
		// Do not check descriptions or short names
		if strings.Contains(key, "_desc") || strings.Contains(key, "_short") {
			continue
		}
		if collectionsByName[strings.ToLower(key)] == nil {
			log.Fatalf("Error parsing collections: %s (%s) was not found in resources/collections.json, make sure to add it manually.", key, value)
		}
	}
}

func parseVdf(lines []string) map[string]string {
	vdf := map[string]string{}

	spacePattern := regexp.MustCompile(`"\s+"`)

	for _, line := range lines {
		// Ignore empty lines or comments
		if len(line) == 0 || line[0:2] == "//" {
			continue
		}
		lineWithoutSpaces := spacePattern.ReplaceAllString(line, `""`)

		parts := strings.Split(lineWithoutSpaces, `""`)
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
	startSaving := false

	r := bufio.NewReader(f)
	bytes, _, err := r.ReadLine()
	for err == nil {
		line := strings.TrimSpace(string(bytes[:]))
		if line == start {
			startSaving = true
		}
		if startSaving {
			if line == end {
				return lines
			}

			lines = append(lines, line)
		}
		bytes, _, err = r.ReadLine()
	}

	return lines
}
