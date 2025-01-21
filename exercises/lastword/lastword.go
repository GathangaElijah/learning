package main

import "fmt"

func LastWord(s string) string {
	var allWords []string
	var word = ""
	strlen := len(s)
	if len(s) == 0 {
		return ""
	}
	for i := 0; i < strlen; i++ {
		if string(s[i]) == " " {
			if word != ""{
				allWords = append(allWords, word)
				word = ""
			}
		} else {
			word += string(s[i])
			if i == len(s)-1 && word != ""{
				allWords = append(allWords, word)
			}
			
		}

	}
	lastword := allWords[len(allWords)-1]
	return lastword
}

func main() {
	fmt.Println(LastWord(" The    ... quick brown fox jumped over the lazy dog"))
	fmt.Println(LastWord(" The quick brown .... fox   jumped over the lazy "))
	fmt.Println(LastWord("   The.quick.brown.fox    "))
	fmt.Println(" ")
}
