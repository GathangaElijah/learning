package main

import "fmt"


func StringCount(str string)int{
    var result int
    word := ""
    resultSlice := []string{}
    
    for i := 0; i < len(str); i++{
        char := str[i]
        if string(char) != " " {
			word += string(char)
        } else if string(char) == " " {
			resultSlice = append(resultSlice, word)
			word = ""
		}
    }
	fmt.Println("This is the resultSlice b4 the last word ", resultSlice)
	if word != ""{
		resultSlice = append(resultSlice, word)
	}
    result = len(resultSlice)
    return result
}

func main(){
str := "this is a string with a few words"
result := StringCount(str)
fmt.Println(result)
}