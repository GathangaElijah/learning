package asciiart

import (
	"errors"
	"fmt"
	// "log"
	"strings"
)

func AsciiProcessor(txtInput string, asciiFile []byte) (string, error) {
	// fmt.Printf("Ascii file => \n %v", string(asciiFile))
	var output strings.Builder
	if txtInput == "" {
		output.WriteString("")
		return output.String(), errors.New("empty string")
	} else if txtInput == "\n" {
		output.WriteString("\n")
		return output.String(), nil
	} else {
		asciiChars := strings.Split(string(asciiFile), "\n")
	inputSubStr := strings.Split(txtInput, "\\n")

	for _, txt := range inputSubStr {
		if txt == "\n" || txt == "" {
			fmt.Println()
			continue
		}
		for i := 0; i < 8; i++ {
			for _, txtChar := range txt {
				start := int(((txtChar - 32) * 9) + 1)
				output.WriteString(asciiChars[start+i])
			}
		}
	}

	}
	
	// Create a map of runes between 32 and 126 as keys and
	// the ascii representation 0f 9 lines as the values.
	// asciiMap := make(map[rune][]string)
	// runes := []rune{32-126}
	// asciiValue := asciiChars[:10]
	// asciiMap[runes] = asciiValue
	return output.String(), nil
}
