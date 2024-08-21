package main

import "fmt"

func ThirdTimeIsACharm(str string) string{
	if str == ""{
		return "\n"
	} else if len(str) < 3 {
		return "\n"
	}
	count := 0
	result := ""
	for _, char := range str{
		count ++
		if count == 3{
			result += string(char)
			count = 0
		}
	}
	return result + "\n"
}

func main() {
	fmt.Print(ThirdTimeIsACharm("123456789"))
	fmt.Print(ThirdTimeIsACharm(""))
	fmt.Print(ThirdTimeIsACharm("a b c d e f"))
	fmt.Print(ThirdTimeIsACharm("12"))
}
