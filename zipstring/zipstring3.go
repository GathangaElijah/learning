package main

import "strconv"


func Zipstring3(s string)string{
	result := ""
	count := 1
	for i := 0; i < len(s); i++{
		char := string(s[i])
		if i > 0 && char == string(s[i-1]){
			count ++
			
		}else {
			result += strconv.Itoa(count) + char
			count = 1
		}
	
	}
	result += strconv.Itoa(count) + string(s[len(s)-1])
	return result
}

