package asciiart

import (
	"errors"
	"fmt"
	"strings"
)

func AsciiProcessor(txtInput string, asciiFile []byte) (string, error) {
	var output string
	if txtInput == "" {
		return "", errors.New("empty input")
	} else if txtInput == "\n" {
		return "\n", nil
	} else {
		asciiChars := strings.Split(string(asciiFile), "\n")
		inputSubStr := strings.Split(txtInput, "\n")
		for _, txt := range inputSubStr {
			if string(txt) == "\n" {
				fmt.Println()
				continue
			}
			for _, txtChar := range txt {
				for i := 0; i < 8; i++ {
					for _, asciiChar := range asciiChars {
						start := int((txtChar-32*9)+1)
						
					}
				}
			} 
		}
	}

	return output, nil
}
