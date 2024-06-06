package statistics

import "math"

// This function uses variance to calculate the standard deviation.
func StdDev(allData []float64) float64 {
	var standardDev float64
	standardDev = math.Sqrt(Variance(allData))
	return standardDev
}
