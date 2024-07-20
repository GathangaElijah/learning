package main

import "fmt"

func LastWord(s string)string{

}

func main(){
	fmt.Println(LastWord("The quick brown fox jumped over the lazy dog"))
	fmt.Println(LastWord(" The quick brown .... fox   jumped over the lazy "))
	fmt.Println(LastWord("   The.quick.brown.fox    "))
	fmt.Println(" ")
}