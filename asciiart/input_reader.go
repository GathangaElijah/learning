package asciiart

import (
	"errors"
	"os"
)

func InputReader(strArgs []string) (string, error) {
	if len(strArgs) != 1 {
		return "", errors.New("the program requires a single string input")
	}
	argOutput := os.Args[1]
	return argOutput, nil
}
