package main

import (
	"log"
	"os"

	"github.com/EarvinKayonga/tasks/commands"
)

func main() {
	app := commands.Create()
	err := app.Run(os.Args)
	if err != nil {
		log.Fatal(err)
	}
}
