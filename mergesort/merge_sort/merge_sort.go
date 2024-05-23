package merge_sort

import (
	"fmt"
)

// Receives a slice of float64 and returns a sorted slice of float64
var dataSet = []float64{}
var midpoint = len(dataSet) / 2
var	firstHalve = dataSet[:midpoint]
var	secondHalve = dataSet[midpoint:]
func MergeSort(firstHalve, secondHalve []float64) ([]float64, []float64){
	// step 1: Divide the dataset  into 2 equal halves
	
	for len(firstHalve) != 1 && len(secondHalve) != 1 {
		firstHalve, secondHalve = MergeSort(firstHalve,secondHalve)
	}
	fmt.Println("This is my firsthalve => ", firstHalve)
	fmt.Println("This is my secondhalve => ", secondHalve)
	return firstHalve, secondHalve
}
