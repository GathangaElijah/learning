package main

import (
	"log"
	"os"

	"learning/asciiart/asciiart"
)

func main() {
	// Sets how we view the errors from reading the commandline arguments
	// ignoring the date and time.
	log.SetPrefix(" Input Error: ")
	log.SetFlags(0)

	args := os.Args[1:]
	output, err := asciiart.InputReader(args)
	if err != nil {
		log.Fatal(err)
	}
	asciiart.AsciiProcessor(output)
}
