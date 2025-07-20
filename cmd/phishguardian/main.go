// cmd/phishguardian/main.go
package main

import (
	"flag"
	"log"
	"os"

	"phishguardian/internal/phishguardian"
)

func main() {
	verbose := flag.Bool("verbose", false, "Enable verbose logging")
	flag.Parse()

	app := phishguardian.NewApp(*verbose)
	if err := app.Run(); err != nil {
		log.Fatal(err)
	}
}
