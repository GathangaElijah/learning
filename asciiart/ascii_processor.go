package asciiart

import (
	"errors"
	"fmt"
	"strings"
)

func AsciiProcessor(txtInput string, asciiFile []byte) (string, error) {
	if txtInput == "" {
		return "", errors.New("empty string")
	} else if txtInput == "\\n" {
		return "\n", nil
	}
	asciiChars := strings.Split(string(asciiFile), "\n")
	inputSubStr := strings.Split(txtInput, "\\n")

	for _, txt := range inputSubStr {
		if txt == "\\n" || txt == "" {
			fmt.Println()
			continue
		}
		for i := 0; i < 8; i++ {
			for _, txtChar := range txt {
				start := int(((txtChar - 32) * 9) + 1)
				fmt.Print(asciiChars[start+i])
			}
			fmt.Println()
		}
	}
	return "", nil
}
