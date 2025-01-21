package main

import "fmt"

func main() {
	fmt.Print(FifthAndSkip("abcdefghijklmnopqrstuwxyz"))
	fmt.Print(FifthAndSkip("This is a short sentence"))
	fmt.Print(FifthAndSkip("1234"))
}

func FifthAndSkip(str string) string {
	var finalStr string

	if len(str) < 5 {
		return "Invalid Input" + "\n"
	}
	if str == "" {
		return "\n"
	}

	word := ""
	for i, char := range str {
		if char != ' ' {
			word += string(char)
			if len(word) == 6 {
				finalStr += word[:len(word)-1] + " "
				word = ""
			}
		} else {
			continue
		}
		if i == len(str)-1 && word != "" {
			finalStr += word
		}
	}

	return finalStr + "\n"
}
