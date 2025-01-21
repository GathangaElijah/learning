package main

func Contains(s, sub string)bool{
	sLen := len(s)
	subLen := len(sub)

	if subLen > sLen{
		return false
	}
	for i := 0; i <= sLen-subLen; i++{
		if s[i: i+sLen] == sub{
			return true
		}
	}
	
	return false
}

