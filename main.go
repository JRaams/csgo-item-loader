package main

import (
	"fmt"
	"log"
	"os"

	"github.com/urfave/cli/v2"
)

func main() {
	app := &cli.App{
		Name:        "csgo-item-loader",
		Description: "Parses CS:GO items from game files into api-ready json.",
		Commands: []*cli.Command{
			{
				Name:        "download",
				Aliases:     []string{"d"},
				Description: "download game files",
				Action: func(c *cli.Context) error {
					fmt.Println("Downloading game files!")
					return nil
				},
			},
			{
				Name:        "parse",
				Aliases:     []string{"p"},
				Description: "parse game files into txt files",
				Action: func(c *cli.Context) error {
					fmt.Println("Parsing game files!")
					return nil
				},
			},
			{
				Name:        "stitch",
				Aliases:     []string{"s"},
				Description: "stitch data into json",
				Action: func(c *cli.Context) error {
					fmt.Println("Stitching into json")
					return nil
				},
			},
		},
	}

	err := app.Run(os.Args)
	if err != nil {
		log.Fatal(err)
	}
}
