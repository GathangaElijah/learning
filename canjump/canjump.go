package main

import "fmt"

func CanJump(s []uint) bool{
	if len(s) == 0 {
		return false
	} else if len(s) == 1{
		return true
	} 
	position := 0

	for position <= len(s){
		steps := int(s[position])

		position += steps

		if position == len(s)-1{
			return true
		}

		if position >= len(s) || steps == 0{
			return false
		}
	}
	return false
}

func main() {
	input1 := []uint{2, 3, 1, 1, 4}
	fmt.Println(CanJump(input1))

	input2 := []uint{3, 2, 1, 0, 4}
	fmt.Println(CanJump(input2))

	input3 := []uint{0}
	fmt.Println(CanJump(input3))
}
