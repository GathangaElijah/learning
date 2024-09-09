package main

import (
	"fmt"
	"os"
)
func Replace(s, old, new string) string{
	var newStr string
	i:= 0
	for i < len(s){
		if i + len(old) <= len(s) && s[i:i+len(old)] == old{
			newStr += new
			i += len(old)
		}else {
			newStr += string(s[i])
			i++
		}
	}
	return newStr
}

func main(){
	args := os.Args
	if len(args) != 4{
		// fmt.Println("\n")
		return 
	}
	str := args[1]
	old := args[2]
	new := args[3]

	result := Replace(str,old,new)
	
	fmt.Println(result)
}