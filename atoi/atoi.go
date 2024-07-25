package main

import "fmt"

func Atoi(s string) int {
var result int = 0
started := false
sign := 1
for _, char := range s{
	if char == ' ' && !started{
		continue
	}
	if char == '-' && !started{
		started = true
		sign = -1
		continue
	}
	if char == '+' && !started{
		started = true
		continue
	}
	if char >= '0' && char <= '9'{
		digitval := int(char -'0')
		result = result * 10 + digitval
		started = true
	} else if char == ' ' && started{
		return 0
	} else {
		break
	}
}

return result *sign
}

func main(){
	fmt.Println(Atoi("12345"))
	fmt.Println(Atoi("0000000012345"))
	fmt.Println(Atoi("012 345"))
	fmt.Println(Atoi("Hello World!"))
	fmt.Println(Atoi("+1234"))
	fmt.Println(Atoi("-1234"))
	fmt.Println(Atoi("++1234"))
	fmt.Println(Atoi("--1234"))
}