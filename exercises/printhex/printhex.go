package main

import (
	"fmt"
	"os"
)

func atoa(str string)int{
result := 0
started := false
sign := 1
for _, char := range str{
	if char == ' ' && !started{
		continue
	}
	if char == '-' && !started{
		sign = -1
		started = true
		continue
	}
	if char == '+' && !started{
		started = true
		continue
	}
	if char >= '0' && char <= '9'{
		digitval := int(char - '0')
		result = result * 10 + digitval
		started = true
	}else {
		break
	}

}
return result *sign
}

func main(){
	args := os.Args
	if len(args) != 2 {
		return
	}
	str := args[1]
	for _, char := range str{
		if char == ' '{
			fmt.Println("ERROR")
			return
		}
	}
	num := atoa(str)
	result := ""
	for num > 0 {
		remainder := num%16
		num /= 16

		var hexChar  byte
		if remainder < 10 {
			hexChar = byte('0' + remainder)
		}else {
			hexChar = byte('a' + (remainder-10) )
		}
		result = string(hexChar) + result
	}
	fmt.Println(result)
}