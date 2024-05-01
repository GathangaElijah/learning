package asciiart

import (
	"errors"
)

func InputReader(strArgs []string) (string, error) {
	if len(strArgs) != 1 {
		return "", errors.New("empty input")
	}
	return "", nil
}
