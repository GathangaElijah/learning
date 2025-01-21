package main

import "fmt"

func Split(str, delim string) []string{
	var result []string
	var word string
	delimLen := len(delim)

	for i := 0; i < len(str);{
		if len(str[i:]) >= delimLen && str[i:i+delimLen] == delim{
			result = append(result, word)
			word = ""
			i += delimLen
		}else {
			word += string(str[i])
			i ++
		}
	}
	result = append(result, word)

	return result
}


func ExpandString(s string) string{
	var expandedString string

	slices := Split(s, " ")
	for i, word := range slices{
		if i != len(slices)-1{
			expandedString += word + "   "
		}else {
			expandedString += word
		}
	}

	return expandedString
}

func main(){
	fmt.Println(ExpandString("The quick brown fox"))
}