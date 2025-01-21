package main

import (
	"strconv"
)

func ZipString2(s string) string{
	result := ""
	words := []rune(s)
	i := 0

	for i < len(s){
		if words[i] == ' '{
			result += " "
			i++
			continue
		}

		var seenChar = make(map[rune]bool)
		for i < len(s) && words[i] != ' '{
			char := words[i]
			if seenChar[char]{
				i++
				continue
			}
			count := 1
			for j := i+1; j < len(s) && words[j] != ' '; j++{
				if words[j] == char{
					count ++
				}
			}
			result += strconv.Itoa(count) + string(char)
			seenChar[char] = true
			i++
		}
	}
	return result
}
	
	
	