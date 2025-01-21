package main

import "fmt"

func WeAreUnique(str1, str2 string)int{
	if str1 == "" && str2 == ""{
		return -1
	}
	unique1 := make(map[rune]bool)
	unique2 := make(map[rune]bool)

	for _, char := range str1{
		unique1[char] = true
	}

	for _, char := range str2{
		unique2[char] = true
	}

	count := 0
	 for char := range unique1 {
		if !unique2[char]{
			count ++
		}
	 }
	 for char := range unique2 {
		if !unique1[char]{
			count++
		}
	 }
	return count
}

func main() {
	fmt.Println(WeAreUnique("foo", "boo"))
	fmt.Println(WeAreUnique("", ""))
	fmt.Println(WeAreUnique("abc", "def"))
}
