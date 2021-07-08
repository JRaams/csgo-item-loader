package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"strconv"
	"strings"
)

func parse() {
	// fmt.Println("Parsing collections files...")
	// _ = parseCollections()
	// fmt.Println("resources/collections.json is up to date!")

	fmt.Println("Parsing weapons")
	weapons := parseWeapons()
	fmt.Println(weapons)
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

// Check if all actual collections have an entry in resources/collections
// If not, it needs to be added manually
func validateCollections() map[string]*Collection {
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
	// Remove first 2 and last line
	text = text[2:len(text)-1]

	// 	"1"
	// 	{
	// 		"name"		"weapon_deagle"
	// 		"prefab"		"weapon_deagle_prefab"
	// 		"item_quality"		"normal"
	// 		"baseitem"		"1"
	// 		"default_slot_item"		"1"
	// 		"item_sub_position"		"secondary4"
	// 		"item_shares_equip_slot"		"1"
	// 	}

	weaponsById := map[int]*Weapon{}

	var id int;
	var err error;
	vdfLines := []string{};
	for _, line := range text {
		if strings.HasPrefix(line, "		\"") {
			id, err = strconv.Atoi(strings.ReplaceAll(strings.ReplaceAll(line, "\"", ""), "\t", ""))
			if err != nil {
				continue;
			}
			print(id);
			vdfLines = vdfLines[:0];
			continue;
		}
		if line == "		{" {
			continue;
		}
		if line == "		}" {
			if id == 0 {
				continue;
			}
			parsedVdf := parseVdf(vdfLines);
			w := Weapon{
				ID: id,
				Name: parsedVdf["name"],
			}
			weaponsById[id] = &w;
			print(parsedVdf);
		}
		vdfLines = append(vdfLines, line);

		print(id);
	}

	print(weaponsById);

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
