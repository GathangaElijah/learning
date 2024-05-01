package asciiart

func AsciiProcessor(txtInput string) (string, error) {
	if txtInput == "\n" {
		return "\n", nil
	}
	return txtInput, nil
}
