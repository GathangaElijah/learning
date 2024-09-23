package main

import (
	"fmt"
	"learning/asciiart/mergesort/merge_sort"
)

func main() {
	var dataSet = []float64{78, 10, 8, 99, 65, 3, 5, 17, 15, 31,33}
	firstHalve, secondHalve := merge_sort.MergeSort(dataSet)
	fmt.Println("This is the sorted dataset =>", firstHalve, secondHalve)
}
