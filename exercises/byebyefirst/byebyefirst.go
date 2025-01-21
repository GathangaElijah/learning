package main

import "fmt"

func ByeByeFirst(strings []string)[]string{
	if len(strings) == 0 {
		return nil
	}
	result := []string{}
	result = append(result, strings[1:]...)
	return result
}

func main() {
	fmt.Println(ByeByeFirst([]string{}))
	fmt.Println(ByeByeFirst([]string{"one arg"}))
	fmt.Println(ByeByeFirst([]string{"first", "second"}))
	fmt.Println(ByeByeFirst([]string{"", "abcd", "efg"}))
}
