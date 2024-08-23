package main

import "fmt"

func RevConcatAlternate(slice1, slice2 []int)[]int{
result := []int{}
// i:= 0
len1, len2 := len(slice1), len(slice2)
if len1 >= len2{
	for i := len1-1; i >= len2; i--{
		result = append(result, slice1[i])
	}
	for i := len2-1; i >= 0; i--{
		result = append(result, slice1[i], slice2[i])
	}
} else {
	for i := len2-1; i >= len1; i--{
		result = append(result, slice2[i])
	}
	for i := len1-1; i >= 0; i--{
		result = append(result, slice1[i], slice2[i])
	}
}
return result
}

func main() {
	fmt.Println(RevConcatAlternate([]int{1, 2, 3}, []int{4, 5, 6}))
	fmt.Println(RevConcatAlternate([]int{1, 2, 3}, []int{4, 5, 6, 7, 8, 9}))
	fmt.Println(RevConcatAlternate([]int{1, 2, 3, 9, 8}, []int{4, 5}))
	fmt.Println(RevConcatAlternate([]int{1, 2, 3}, []int{}))
}
