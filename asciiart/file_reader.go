package asciiart

import (
	"bufio"
	"log"
	"os"
)

func ShadowReader() []byte {
	shadow_file, err := os.Open("shadow.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer shadow_file.Close()
	file_stat, err := shadow_file.Stat()
	if err != nil {
		log.Fatal(err)
	}
	file_size := file_stat.Size()
	reader := bufio.NewReader(shadow_file)
	var shadow_data = make([]byte, file_size)
	_, err = reader.Read(shadow_data)
	if err != nil {
		log.Fatal(err)
	}
	return shadow_data
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
	thinkertoy_file, err := os.Open("thinkertoy.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer thinkertoy_file.Close()
	file_stat, err := thinkertoy_file.Stat()
	if err != nil {
		log.Fatal(err)
	}
	file_size := file_stat.Size()
	var thinkertoyData = make([]byte, file_size)
	reader := bufio.NewReader(thinkertoy_file)
	_, err = reader.Read(thinkertoyData)
	if err != nil {
		log.Fatal()
	}
	return thinkertoyData
}
