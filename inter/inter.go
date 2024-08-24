package main

import (
	"fmt"
	"os"
)

func main(){
	args := os.Args
	if len(args) != 3{
		return
	}
	str1 := args[1]
	str2 := args[2]

	result := ""
	charMap := make(map[rune]bool)
	for _, char := range str2{
		charMap[char] = true
	}
	seen := make(map[rune]bool)
	for _, char := range str1{
		if charMap[char] && !seen[char]{
			result += string(char)
			seen[char] = true
		}
	}
	fmt.Println(result)
}