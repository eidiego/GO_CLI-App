package main

import (
	"cli_server/app"
	"fmt"
	"log"
	"os"
)

func main() {
	fmt.Println("Starting main function...")

	application := app.Generate()

	if error := application.Run(os.Args); error != nil {
		log.Fatal(error)
	}
}
