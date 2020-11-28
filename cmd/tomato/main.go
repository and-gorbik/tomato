package main

import (
	"log"
	"os"

	"tomato/internal/cli"
)

func main() {
	app := cli.New()
	if err := app.Run(os.Args); err != nil {
		log.Fatal(err)
	}
}
