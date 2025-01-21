package main

import "fmt"

func ConcatAlternate(slice1, slice2 []int) []int {
	len1 := len(slice1)
	len2 := len(slice2)

	var finalSlice []int
	if len2 > len1{
		for i := 0; i < len1; i++{
			finalSlice = append(finalSlice, slice2[i], slice1[i])
		}
		finalSlice = append(finalSlice, slice2[len1:]...)
	}else {
		for i := 0; i < len2; i ++{
			finalSlice = append(finalSlice, slice1[i], slice2[i])
		}
		finalSlice = append(finalSlice, slice1[len2:]...)
	}

	return finalSlice
}

func main() {
	fmt.Println(ConcatAlternate([]int{1, 2, 3}, []int{4, 5, 6}))
	fmt.Println(ConcatAlternate([]int{2, 4, 6, 8, 10}, []int{1, 3, 5, 7, 9, 11}))
	fmt.Println(ConcatAlternate([]int{1, 2, 3}, []int{4, 5, 6, 7, 8, 9}))
	fmt.Println(ConcatAlternate([]int{1, 2, 3}, []int{}))
}
