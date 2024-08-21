package main

import (
	"fmt"
	"os"
)

func main(){
	args := os.Args
	if len(args) != 4{
		// fmt.Println("\n")
		return 
	}
	str1 := args[1]
	str2 := args[2]
	str3 := args[3]

	result := ""
	for _, char := range str1{
		if char == rune(str2[0]){
			result += str3
		} else{
			result += string(char)
		}
	}
	fmt.Println(result)
}