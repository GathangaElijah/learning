package asciiart

import (
	"flag"
	"errors"
	"learning/asciiart/data"
	"strings"
)

func AsciiProcessor(txtInput string, file *data.AsciiFiles) (string, error) {
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
	var asciiFlag string
	flag.StringVar(&asciiFlag, "type", "", "Use st (Standard), th(thinkertoy, sh(shadow))")
	flag.Parse()

	switch asciiFlag{
	case "th":
		asciiChars = strings.Split(string(file.Thinkertoy), "\r\n")
	case "st":
		asciiChars = strings.Split(string(file.Standard), "\n")
	case "sh":
		asciiChars = strings.Split(string(file.Shadow), "\n")
	default:
		return "", errors.New("invalid flag type")
	}
	
	// var standardFlag = flag.String("st", "", "use st for standard ascii")
	// 	var shadowFlag = flag.String("sh", "", "use sh for shadow file")

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
