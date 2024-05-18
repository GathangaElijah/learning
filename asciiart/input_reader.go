package asciiart

import (
	"errors"
	"os"
)

func InputReader(strArgs []string) (string, error) {
	if len(strArgs) != 3 {
		return "", errors.New("the program requires a single string input")
	}
	argOutput := os.Args[3]
	return argOutput, nil
}
