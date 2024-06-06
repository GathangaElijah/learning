package statistics

import "math"

func Variance(allData []float64) float64 {
	mean := Average(allData)

	var variance float64
	n := float64(len(allData))
	var sqDiff float64
	for _, num := range allData {
		diff := num - mean
		sqDiff += math.Pow(diff, 2)
	}
	variance = sqDiff / n
	return variance
}
