package main

import "strconv"


func ZipString3(s string) (str string) {
	count := 1

	for i := 0; i < len(s); i++ {
		if i > 0 && s[i-1] == s[i] {
			count++
			str = str[:len(str)-2] + strconv.Itoa(count) + string(s[i])
		} else {
			count = 1
			str += strconv.Itoa(count) + string(s[i])
		}
	}
	return
}

