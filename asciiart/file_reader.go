package asciiart

import (
	"log"
	"os"
)

func ShadowReader() []byte {
	shadowfile, err := os.ReadFile("shadow.txt")
	if err != nil {
		log.Fatal(err)
	}
	return shadowfile
}

func StandardReader() []byte {
	standard_file, err := os.ReadFile("standard.txt")
	if err != nil {
		log.Fatal(err)
	}
	return standard_file
}

func ThinkertoyReader() []byte {
	thinkertoy_file, err := os.ReadFile("thinkertoy.txt")
	if err != nil {
		log.Fatal(err)
	}
	return thinkertoy_file
}
