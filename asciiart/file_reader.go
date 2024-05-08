package asciiart

import (
	"bufio"
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
	standard_file, err := os.Open("standard.txt")
	if err != nil {
		log.Fatal(err)
		return nil
	}
	defer standard_file.Close()
	file_stat, err := standard_file.Stat()
	if err != nil {
		log.Fatal(err)
	}
	fileSize := file_stat.Size()
	reader := bufio.NewReader(standard_file)
	var standardData = make([]byte, fileSize)
	_, err = reader.Read(standardData)
	if err != nil {
		log.Fatal(err)
		return nil
	}
	return standardData
}

func ThinkertoyReader() []byte {
	thinkertoy_file, err := os.ReadFile("thinkertoy.txt")
	if err != nil {
		log.Fatal(err)
	}
	return thinkertoy_file
}
