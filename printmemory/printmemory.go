package main

import "fmt"

func main(){
	fmt.Println(PrintMemory([10]byte{'h', 'e', 'l', 'l', 'o', 16, 21, '*'}))

}

func toHex(n int) string{
	hexVal := "0123456789ABCDEF"
	var result []string
	var finalString string
	for n > 0 {
		remainder := n%16
		result = append(result, string(hexVal[remainder]))
		n/=16
	}
	// Ensure at least two characters are returned (e.g., "0A" instead of "A")
	if len(result) == 1 {
		result = append(result, "0")
	}
	if len(result) == 0 {
		result = append(result, "0", "0")
	}
	// reverse the string
	for i, j := 0, len(result)-1; i < j; i, j = i+1, j-1{
		result[i], result[j] = result[j], result[i]
	}

	for _, char := range result{
		finalString += char + ""
	}

	return finalString
}

func PrintMemory(arr [10]byte)string{
var output string
var memory []string
str := ""
for _, char := range arr{
	memory = append(memory, toHex(int(char)))
	if char >= 32 && char <= 126{
		str += string(char)
	} else {
		str += "."
	}
}

// fmt.Printf("Memory %v length %v", memory, len(memory))
for i := 0; i < len(memory); i+=4{
	hexpart := ""
	for j := i; j < i+4 && j < len(memory); j++{
		hexpart += memory[j] + " "
	}
	output += hexpart + "\n" 
}

// fmt.Println(str)
return output  + str
}
