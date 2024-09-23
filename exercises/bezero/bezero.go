package main

import "fmt"

func BeZero(s []int)[]int{
	if len(s) == 0 {
		return nil
	}
	result := []int{}
	for i := 0; i < len(s); i++{
		s[i] = 0
		result = append(result, s[i])
	}
	return result
}

func main() {
	fmt.Println(BeZero([]int{1, 2, 3, 4, 5, 6}))
	fmt.Println(BeZero([]int{}))
}
