package main

import (
	"fmt"
	"strconv"
)

func main() {
	val1 := true
	val2 := false

	str1 := strconv.FormatBool(val1)
	str2 := strconv.FormatBool(val2)

	fmt.Println("Formatted value 1: " + str1)
	fmt.Println("Formatted value 2: " + str2)
}

/*
In this chapter,
1. I have learnt the basic data types how to use them.
2. I have learn data conversion by using explicit methods and using the strConv package
to convert integer values to strings, float values to strings, string values to integers
or to floats and boolean values to string values.
3. Formatting of values to strings.
4. I have also learnt about the operations using 
a) mathematical operators for maths related operations.
b) comparison operators 
c) logical operators
d) assignment operators
e) increament and decreament operators

*/
