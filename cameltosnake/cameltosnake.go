package main

import (
	"fmt"
)
func  isLetter(s rune) bool{
	return(s >= 'A' && s <= 'Z') || (s >= 'a' && s <= 'z')
	
}

func isUpper(s rune) bool{
	return s >= 'A' && s <= 'Z'
	
}
func toLower(s rune)rune{
	if isUpper(s){
		return s + ('a' - 'A')
	}
	return s
}

func isCamelCase(s string) bool{
	if s == ""{
		return false
	}
	
	for i := 0; i < len(s); i++{
		char := rune(s[i])
		if i > 0 && isUpper(char) && isUpper(rune(s[i-1])){
			return false
		} 
		if isUpper(rune(s[len(s)-1])){
			return false
		} 
		if !isLetter(rune(s[i])){
			return false
		}
	}
	return true
}

func CamelToSnakeCase(s string) string {
	if s == ""{
		return ""
	} else if !isCamelCase(s){
		return s
	}
	result := ""

		for i := 0; i < len(s); i++{
			char := rune(s[i])
			if isUpper(char) && i > 0{
				result +="_"
			}
			result += string(char)
		}
	return result
}

func main() {
	fmt.Println(CamelToSnakeCase("HelloWorld"))
	fmt.Println(CamelToSnakeCase("helloWorld"))
	fmt.Println(CamelToSnakeCase("camelCase"))
	fmt.Println(CamelToSnakeCase("CAMELtoSnackCASE"))
	fmt.Println(CamelToSnakeCase("camelToSnakeCase"))
	fmt.Println(CamelToSnakeCase("hey2"))
}
