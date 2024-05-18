package main

import (
	"fmt"
	"log"
	"os"

	"learning/asciiart/asciiart"
	"learning/asciiart/data"
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
	thinkertoyData := asciiart.ThinkertoyReader()
	shadowData := asciiart.ShadowReader()
	standardData := asciiart.StandardReader()
	// This part creates an instance of the AsciiFiles struct
	// which is in the data folder.
	file := &data.AsciiFiles{
		Thinkertoy: thinkertoyData,
		Standard:   standardData,
		Shadow:     shadowData,
	}

	asciiRep, err := asciiart.AsciiProcessor(strInput, file)
	log.SetPrefix("Processing Error: ")
	if err != nil {
		// return
		log.Fatal(err)
	}
	fmt.Print(asciiRep)
}
