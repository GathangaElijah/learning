package main

import "fmt"

func AtoiBase(s string, base string) int {
	if !isValid(base){
		return 0
	}
	result := 0
	length := len(base)
	for _, sChar := range s{
		index := indexOf(sChar, base)
		if index == -1{
			return 0
		}
		result = result*length + index
	}
	return result
}

func isValid(base string)bool{
	if len(base) < 2{
		return false
	}
	var charMap = make(map[rune]bool)
	for _, char := range base{
		if char == '+' || char == '-'{
			return false
		} else {
			charMap[char] = true
		}
	}
	return true
}

func indexOf(char rune, base string) int{
	for i, bChar := range base{
		if char == bChar{
			return i
		}
	}
	return -1
}

func main() {
	fmt.Println(AtoiBase("125", "0123456789"))
	fmt.Println(AtoiBase("1111101", "01"))
	fmt.Println(AtoiBase("7D", "0123456789ABCDEF"))
	fmt.Println(AtoiBase("uoi", "choumi"))
	fmt.Println(AtoiBase("bbbbbab", "-ab"))
}
