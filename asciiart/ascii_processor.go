package asciiart

import "errors"

func AsciiProcessor(txtInput string) (string, error) {
	if txtInput == "" {
		return "", errors.New("empty string")
	} else if txtInput == "\n" {
		return "\n", nil
	}

	return txtInput, nil
}
