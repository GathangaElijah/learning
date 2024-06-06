package statistics

import (
	"sort"
)

/*
Here we reuse the allData variable from the StringConverter()
function in average.go file. It will be use by the Median() function
to find the median of the data.
*/
func Median(allData []float64) float64 {
	var median float64
	copiedData := make([]float64, len(allData))
	copy(copiedData, allData)
	sort.Float64s(copiedData)
	var n int = len(copiedData)
	if n%2 == 0 {
		median = (copiedData[(n/2)-1] + copiedData[(n/2)]) / 2
	} else {
		median = copiedData[n/2]
	}

	return median
}
