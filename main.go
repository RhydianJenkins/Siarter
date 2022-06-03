package main

import (
	"fmt"
	"os"
	"time"

	"github.com/urfave/cli/v2"
)

var (
	cliApp *cli.App
)

func init() {
	cliApp = &cli.App{
		Name:        "siarter",
		Version:     "v0.0.1",
		Description: "A AIS charter TUI",
		Compiled:    time.Now(),
		Authors: []*cli.Author{
			{
				Name:  "Rhydian",
				Email: "rhydz@msn.com",
			},
		},
		Commands: []*cli.Command{
			{
				Name:    "chart",
				Aliases: []string{"c"},
				Usage:   "Show a boat's current AIS information",
				Action: func(c *cli.Context) error {
					return nil
				},
			},
		},
		Flags: []cli.Flag{},
	}
}

func main() {
	err := cliApp.Run(os.Args)
	if err != nil {
		fmt.Println(err)
	}
}
