package main

import (
	"github.com/urfave/cli/v2"
	"log"
	"os"
	"supply-chain-mall/pkg/core"
)

func main() {
	app := &cli.App{
		Name:  "app",
		Usage: "app",
		Action: func(c *cli.Context) error {
			App, err := NewInjector()
			if err != nil {
				log.Fatalf("wire injector failed: %v", err)
			}
			App.RegisterRoutes()
			if err := core.Run(c.Context, App); err != nil {
				log.Println(err)
			}
			return nil
		},
	}

	if err := app.Run(os.Args); err != nil {
		log.Fatal(err)
	}
}
