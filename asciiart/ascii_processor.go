package asciiart

import (
	"errors"
	"strings"
)

type asciiFiles struct {
	thinkertoy, standard, shadow []byte
}

func AsciiProcessor(txtInput string, file *asciiFiles) (string, error) {
	var output = ""
	if txtInput == "" {
		return "", errors.New("empty string")
	} else if txtInput == "\\n" {
		return "\n", nil
	}

	if file == nil {
		return "", errors.New("nil pointer")
	}

	var asciiChars []string
	if file.thinkertoy != nil{
		asciiChars = strings.Split(string(file.thinkertoy), "\n\r")
	} else if file.standard != nil {
		asciiChars = strings.Split(string(file.standard), "\n")
	} else if file.shadow != nil{
		asciiChars = strings.Split(string(file.shadow), "\n")
	}

	inputSubStr := strings.Split(txtInput, "\\n")

	for _, txt := range inputSubStr {
		if txt == "\\n" || txt == "" {
			output += "\n"
			continue
		}
		for i := 0; i < 8; i++ {
			for _, txtChar := range txt {
				start := int(((txtChar - 32) * 9) + 1)
				output += asciiChars[start+i]
			}
			output += "\n"
		}
	}
	return output, nil
}
