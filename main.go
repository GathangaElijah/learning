package main

import (
	"fmt"
	"log"
	"os"

	"learning/asciiart/asciiart"
)

func main() {
	// Sets how we view the errors from reading the commandline arguments
	// ignoring the date and time.
	log.SetFlags(0)

	args := os.Args[1:]
	strInput, err := asciiart.InputReader(args)
	log.SetPrefix("Input Error: ")
	if err != nil {
		log.Fatal(err)
	}
	asciiFile := asciiart.StandardReader()
	asciiRep, err := asciiart.AsciiProcessor(strInput, asciiFile )
	log.SetPrefix("Processing Error: ")
	if err != nil {
		return
		// log.Fatal(err)
	}
	fmt.Print(asciiRep)
}
