package asciiart

import (
	"errors"
)

func InputReader(strArgs []string) (string, error) {
	if len(strArgs) != 1 {
		return "", errors.New("the program requires a single string input")
	}
	return "", nil
}
