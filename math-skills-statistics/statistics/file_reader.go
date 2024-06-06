package statistics

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

/*
This function opens a file (data.txt) reads the data present
and returns it.
*/

func FileReader(fileName string) []string {
	dataFile, err := os.Open(fileName)
	if err != nil {
		fmt.Println(err)
		os.Exit(0)
	}
	defer dataFile.Close()
	scanner := bufio.NewScanner(dataFile)
	var fileData []string
	for scanner.Scan() {
		data := scanner.Text()
		fileData = append(fileData, data)
	}
	if err := scanner.Err(); err != nil {
		log.Println(err)
	}

	return fileData
}
