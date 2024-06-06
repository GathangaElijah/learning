package statistics

import (
	"fmt"
	"os"
	"strconv"
)

/*
This file has two functions;
1. one the converts the file from FileReader() and returns a slice of floats.
2. Function Average() that receives the returned slice and finds the average
of the present data. It returns a rounded average of data.
*/
func StringConverter(fileData []string) []float64 {
	var allData []float64
	if len(fileData) == 0 {
		fmt.Println("Empty file")
		os.Exit(0)
	}
	for _, num := range fileData {
		data, err := strconv.ParseFloat(num, 64)
		if err != nil {
			fmt.Println("Error occured while converting  data")
			os.Exit(0)
		}
		allData = append(allData, data)
	}
	return allData
}

// Calculating the average of the data
func Average(allData []float64) float64 {
	var sum float64
	var count float64 = 0
	for _, data := range allData {
		count++
		sum = sum + data
	}
	average := sum / count
	return average
}
