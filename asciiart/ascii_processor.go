package asciiart

import (
	"errors"
	// "log"
	"strings"
)

func AsciiProcessor(txtInput string, asciiFile []byte) (string, error) {
	var output strings.Builder
	if txtInput == "" {
		output.WriteString("")
		return output.String(), errors.New("empty input")
	}
	if txtInput == "\n" {
		output.WriteString("\n")
		return output.String(), nil
	}
	asciiChars := strings.Split(string(asciiFile), "\n")
	
	// inputSubStr := strings.Split(txtInput, "\n")
	for _,txt := range txtInput {
		for i := 0; i < 8; i ++ {
			for _, asciiChar := range asciiChars {
				start := int((txt-32)*9 + 1)
				output.WriteByte(asciiChar[start+i])
			}
		}
	}
	return output.String(), nil
}
