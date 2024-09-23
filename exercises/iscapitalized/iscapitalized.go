package main

import "fmt"

func isCapital(r rune)bool{
	if r >= 'A' && r <= 'Z'{
		return true
	}
	return false
}
func nonLetter(r rune)bool{
	return !((r >= 'a' && r <= 'z') || r >= 'A' && r <= 'Z')
}

func IsCapitalized(s string)bool{
	word := ""
	for i := 0; i < len(s); i++{
		char := s[i]
		if char != ' '{
			word += string(char)
		} else{
			word = ""
			continue
		}
		if isCapital(rune(word[0])) || nonLetter(rune(word[0])){
			return true
		}
	}
	return false
}

func main() {
	fmt.Println(IsCapitalized("Hello! How are you?"))
	fmt.Println(IsCapitalized("Hello How Are You"))
	fmt.Println(IsCapitalized("Whats 4this 100K?"))
	fmt.Println(IsCapitalized("Whatsthis4"))
	fmt.Println(IsCapitalized("!!!!Whatsthis4"))
	fmt.Println(IsCapitalized(""))
}
