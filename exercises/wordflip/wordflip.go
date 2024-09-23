package main

import "fmt"

func WordFlip(str string) string {
	// trim the spaces
	trimSpaces := func (s string) string {
		start := 0
		end := len(s)-1
		// finding the first non space character
		for start <= end && s[start] == ' '{
			start ++
		}
		// finding the last non space character
		for end >= start && s[end] == ' '{
			end --
		}
		return s[start:end+1]
	}
	// split string into words
	splitToWords := func (s string) []string  {
		var words []string
		word := ""
		for _, char := range s{
			if char != ' '{
				word += string(char)
			}else if word != ""{
				words = append(words, word)
				word = ""
			}
		}
		if word != ""{
			words = append(words, word)
		}
		return words
	}

	// Step 1:  handle the empty input case
	if len(str) == 0 {
		return "Invalid Output" + "\n"
	}
	//Step 2: trim leading and trailing spaces
	str = trimSpaces(str)
	
	// Step 3: Split into words
	words := splitToWords(str)
	// Step 4 := reverse the order of words
	for i, j := 0, len(words)-1; i < j; i, j = i+1, j-1{
		words[i], words[j] = words[j], words[i]
	}
	// Step 5: Join the words with a single space
	result := ""
	for i, word := range words{
		if i > 0 {
			result += " "
		}
		result += word
	}
	return result + "\n"
}

func main() {
	fmt.Print(WordFlip("First second last"))
	fmt.Print(WordFlip(""))
	fmt.Print(WordFlip("     "))
	fmt.Print(WordFlip(" hello  all  of  you! "))
}
