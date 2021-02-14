package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
)

var toDownload = map[string]string{
	"items_game.txt":     "https://raw.githubusercontent.com/SteamDatabase/GameTracking-CSGO/master/csgo/scripts/items/items_game.txt",
	"items_game_cdn.txt": "https://raw.githubusercontent.com/SteamDatabase/GameTracking-CSGO/master/csgo/scripts/items/items_game_cdn.txt",
	"csgo_english.txt":   "https://raw.githubusercontent.com/SteamDatabase/GameTracking-CSGO/master/csgo/resource/csgo_english.txt",
}

func download() {
	fmt.Printf("Downloading %d game files!\n", len(toDownload))
	// Delete and recreate storage folder
	err := os.RemoveAll("storage")
	if err != nil {
		log.Fatal(err)
	}
	err = os.Mkdir("storage", 0777)
	if err != nil {
		log.Fatal(err)
	}

	// Download all individual files
	for fileName, URL := range toDownload {
		saveFileFromURL(fileName, URL)
	}
}

func saveFileFromURL(fileName string, URL string) {
	fmt.Printf("Downloading '%s'...\n", fileName)
	out, err := os.Create(fmt.Sprintf("storage/%s", fileName))
	if err != nil {
		log.Fatal(err)
	}
	defer out.Close()

	resp, err := http.Get(URL)
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()

	_, err = io.Copy(out, resp.Body)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("Done downloading '%s'!\n", fileName)
}
