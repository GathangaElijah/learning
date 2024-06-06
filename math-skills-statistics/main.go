package main

import (
	"fmt"
	"math"

	"maths/statistics"
)

func main() {
	// Fetching the file
	fileName := statistics.FileFetcher()
	// Reading the file
	lines := statistics.FileReader(fileName)
	// Converting the file
	allData := statistics.StringConverter(lines)
	// Calculating the average
	average := math.Round(statistics.Average(allData))
	fmt.Printf("Average: %.0f\n", average)
	// Calculating the median
	median := math.Round(statistics.Median(allData))
	fmt.Println("Median:", median)
	// Calculating the variance
	variance := math.Round(statistics.Variance(allData))
	fmt.Printf("Variance:%.0f\n", variance)
	// Calculating the standard deviation
	standardDev := math.Round(statistics.StdDev(allData))
	fmt.Printf("Standard Deviation: %.0f\n", standardDev)
}
