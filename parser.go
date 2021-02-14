package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

func parse() {
	fmt.Println("Parsing game files!")
	parseCollections()
}

func parseCollections() {
	text := getTextBetweenLines("storage/csgo_english.txt", "// SET DESCRIPTIONS", "/////////////////")
	fmt.Println(text)
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

			fmt.Println(line)
			lines = append(lines, line)
		}
		bytes, _, err = r.ReadLine()
	}

	return lines
}
