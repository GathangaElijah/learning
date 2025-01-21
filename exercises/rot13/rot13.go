package main

import (
	"fmt"
	"os"
)
/*
This function takes a string and returns a new
string after adding 13.
*/
func Rot13(str string) string {
	fmt.Println("This is my input string =>", str)
	var finalStr []rune
	for _, char := range str {
		if char >= 'a' && char <= 'z' {
			char = 'a' + (char-'a'+13)%26

		}
		finalStr = append(finalStr, char)
	}

	return string(finalStr)
}

func main() {
	str := os.Args[1:]
	if len(str) != 1 {
		fmt.Println("Usage: go run . 'hello'")
		os.Exit(0)
	}
	fmt.Println("This is my output string =>", Rot13(str[0]))
}
