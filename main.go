package main

import (
	"fmt"
	"log"
	"os"
	"time"

	"github.com/joho/godotenv"
	"github.com/rhydianjenkins/siarter/pkg/httpClient"
	"github.com/urfave/cli/v2"
)

var (
	cliApp *cli.App
	apiKey string
	shipId string
	mock   bool
)

func init() {
	err := godotenv.Load(".env")

	if err != nil {
		log.Fatalf("Error loading .env file")
	}

	apiKey = os.Getenv("API_KEY")
	shipId = os.Getenv("SHIP_ID")
	mock = os.Getenv("MOCK_API") == "1"
}

func init() {
	cliApp = &cli.App{
		Name:        "siarter",
		Version:     "v0.0.1",
		Description: "An AIS charter TUI",
		Compiled:    time.Now(),
		Authors: []*cli.Author{
			{
				Name:  "Rhydian",
				Email: "rhydz@msn.com",
			},
		},
		Commands: []*cli.Command{
			{
				Name:    "info",
				Aliases: []string{"i"},
				Usage:   "Show a boat's current AIS information",
				Action: func(c *cli.Context) error {
					url := fmt.Sprintf("https://services.marinetraffic.com/api/exportvessel/%s?v=1&protocol=jsono&shipId=%s", apiKey, shipId)
					client := httpClient.NewClient(url, mock)
					response, err := client.Get()

					if err != nil {
						return err
					}

					fmt.Printf("%+v\n", response) // TODO somehow print this?

					return nil
				},
			},
		},
		Flags: []cli.Flag{},
	}
}

func main() {
	cliApp.Run(os.Args)
}
