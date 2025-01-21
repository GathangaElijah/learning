package main

import "fmt"

func main(){
	// using if statements
	// it executes a group of statements when the specified expression
	//produces a bool value  of true when it is evaluated.

	kayakPrice := 275.00

	if (kayakPrice > 500) {
		fmt.Println("Price is greater than 100")
	} else if (kayakPrice < 100) {
		fmt.Println("Price is less than 100")
	} else {
		fmt.Println("Price not matched by earlier expression")
	}

	fmt.Println()
	fmt.Println("**** FOR LOOPS ****")
	fmt.Println()
	// for keyword creates a loop that repeatedly execute statements
	
}