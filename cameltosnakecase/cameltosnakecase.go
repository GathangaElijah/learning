package main

import "fmt"

func isUpper(ch byte) bool {
	if ch >= 65 && ch <= 90 {
		return true
	}
	return false
}

func isChar(ch byte) bool {
	if ch >= 'A' && ch <= 'Z' || ch >= 'a' && ch < 'z' {
		return true
	}
	return false
}

func camelCase(s string) bool {
	if isUpper(s[len(s)-1]) {
		return false
	}
	for i := 0; i < len(s); i++ {
		if i > 0 && isUpper(s[i]) && isUpper(s[i-1]) {
			return false
		}
		if !isChar(s[i]) {
			return false
		}
	}
	return true
}

func snakeCase(s string) string{
	var str string
	for i := 0; i < len(s); i++{
		if isUpper(s[i]) && i > 0{
			str = s[:i] + "_" + s[i:]
		}
	}
	return str
}

func CamelToSnakeCase(s string) string {
	var finalStr string
	if s == "" {
		return ""
	}

	if !camelCase(s){
		return s
	}
	finalStr = snakeCase(s)
	return finalStr
}

func toUpper(s string) rune {
	var newchar rune
	for _, char := range s {
		if char >= 'a' && char <= 'z' {
			newchar = char - 32
		}
	}
	return newchar
}

func toLower(s string) rune {
	var newchar rune
	for _, char := range s {
		if char >= 'A' && char <= 'Z' {
			newchar = char + 32
		}
	}
	return newchar
}

func main() {
	fmt.Println(CamelToSnakeCase("HelloWorld"))
	fmt.Println(CamelToSnakeCase("helloWorld"))
	fmt.Println(CamelToSnakeCase("camelCase"))
	fmt.Println(CamelToSnakeCase("CAMELtoSnackCASE"))
	fmt.Println(CamelToSnakeCase("camelToSnakeCase"))
	fmt.Println(CamelToSnakeCase("hey2"))
}
