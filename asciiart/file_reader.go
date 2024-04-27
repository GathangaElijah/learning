package asciiart

import (
	"log"
	"os"
)

func ShadowReader() {
	shadowfile, err := os.ReadFile("shadow.txt")
	if err != nil {
		log.Fatal(err)
	}
	os.Stdout.Write(shadowfile)
}

func StandardReader() {
	standard_file, err := os.ReadFile("standard.txt")
	if err != nil {
		log.Fatal(err)
	}
	os.Stdout.Write(standard_file)
}

func ThinkertoyReader() {
	thinkertoy_file, err := os.ReadFile("thinkertoy.txt")
	if err != nil {
		log.Fatal(err)
	}
	os.Stdout.Write(thinkertoy_file)
}
