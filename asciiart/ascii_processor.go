package asciiart

import (
	"bufio"
	"bytes"
	//"errors"
	//"fmt"
	"strings"
)

func AsciiProcessor(txtInput string, asciiFile []byte) (string, error) {
	// if txtInput == "" {
	// 	return "", errors.New("empty input")
	// }
	// if txtInput == "\\n" {
	// 	fmt.Println()
	// }

	asciiChars := strings.Split(string(asciiFile), "\n")
	txtChar := strings.Split(txtInput, "\\n")

	var asciiStr string
	
	for _, inputChar := range txtChar {
		for _, runeChar := range inputChar{
			for i := 0; i < 8; i++ {
				start := int((runeChar-32*9)+1)
				for _, asciiChar := range asciiChars {

					buffer := new(bytes.Buffer)
					asciiRep := bufio.NewWriter(buffer)
					_, err := asciiRep.WriteString(string(asciiChar[start]))
					if err != nil {
						return "", err
					}
					err = asciiRep.Flush()
					if err != nil{
						return "", err
					}
					asciiStr = buffer.String()
			}
		}
	}
	
}
return asciiStr, nil
}