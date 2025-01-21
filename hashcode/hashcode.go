package main

import "fmt"

func HashCode(dec string)string{
	result := ""
	size := len(dec)
	for _, char := range dec{
		newchar := (char + rune(size))%127
		if newchar < 32{
			newchar += 33
		}
		result += string(newchar)
	}
return result
}

func main() {
	fmt.Println(HashCode("A"))
	fmt.Println(HashCode("AB"))
	fmt.Println(HashCode("BAC"))
	fmt.Println(HashCode("Hello World"))
}
