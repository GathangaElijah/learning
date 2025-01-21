package main

import (
	"fmt"
)

func ZipString(s string) string {
	i := 0
	var result string
	runes := []rune(s)
	length := len(runes)

	for i < length {
		currentChar := runes[i]
		if !(currentChar >= 'a' && currentChar <= 'z' || currentChar >= 'A' && currentChar <= 'Z') {
			i++
			continue
		}
		count := 1
		for i+1 < length && runes[i+1] == currentChar {
			count++
			i++
		}
		result += fmt.Sprintf("%d%c", count, currentChar)
		i++
	}

	return result
}

func main() {
	fmt.Println(ZipString3("YouuungFellllas"))
	fmt.Println(ZipString3("Thee quuick browwn fox juumps over the laaazy dog"))
	fmt.Println(ZipString3("Helloo Therre!"))
}
