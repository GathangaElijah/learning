package main

import (
	"fmt"
	"os"
)

func main(){
	args := os.Args
	if len(args) != 2 || args[1] == ""{
		return 
	}
	word := ""
	result := ""
	str := args[1]
	for i := 0; i < len(str); i++{
		char := str[i]
		if char != ' '{
			word += string(char)
		} else if word != ""{
			if result != ""{
				result += "   "
			}
			result += word
			word = ""
		}
	}
	if word != ""{
		if result != ""{
			result += "   "
		}
		result += word
	}
	fmt.Println(result)
}