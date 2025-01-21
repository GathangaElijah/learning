package main

import "fmt"

func AlphaMirror(s string) string {
	result := make([]byte, len(s))
	for i := 0; i < len(s); i++ {
		char := s[i]
		if char >= 'a' && char <= 'z' || char >= 'A' && char < 'Z' {
			var mirror byte
			if char >= 'a' && char <= 'z'{
				mirror = 'a' + ('z'-char)
			} else {
				mirror = 'A' + ('Z'- char)
			}
			result[i] = mirror
		} else {
			result[i] = char
		}
	}
	return string(result)
}

func main() {
	fmt.Println(AlphaMirror("abc"))
	fmt.Println(AlphaMirror("My horse is Amazing."))
}
