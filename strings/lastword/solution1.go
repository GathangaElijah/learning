package main

import "fmt"

func Lastword(s string) string {
	var words []string
	var word = ""
	if len(s) == 0 {
		return ""
	}

	for i := 0; i < len(s); i++{
		if string(s[i]) == " "{
			if word != ""{
				words = append(words, word)
				word = ""
			}
		} else {
			word += string(s[i])
		}
	}
	fmt.Println(words)
	lastword := words[len(words)-1]
	return lastword
}

func main() {
	fmt.Println(Lastword(" The    quick brown fox jumped over the lazy dog"))
	fmt.Println(Lastword("   The quick brown fox ,,,.. over the lazy dog  "))
	// fmt.Println(Lastword("  lazy.dog "))
	// fmt.Println(Lastword("  "))
}
