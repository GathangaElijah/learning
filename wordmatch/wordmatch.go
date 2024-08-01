package main

import "fmt"

func WdMatch(str1, str2 string) string{
	// var charMap = make(map[rune]bool)
	var result = ""
	start := 0
	
		for i := 0; i < len(str1); i++{
			for j:= start; j <len(str2); j++{
				if str1[i] == str2[j]{
					result += string(str1[i])
					start = j
					break
				} 
			}
		}
if result == str1{
	return result
}

return ""
}


func main(){
	string1 := "faya" 
	string2 := "fgvvfdxcacpolhyghbreda"
	fmt.Println(WdMatch(string1, string2))
}