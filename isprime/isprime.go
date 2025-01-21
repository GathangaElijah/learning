package main

func IsPrime(nb int) bool {
	if nb <= 1 {
		return false
	} else if nb <= 3 {
		return true
	} else if nb%2 == 0 || nb%3 == 0 {
		return false
	}
	i := 5
	for i*i <= nb {
		if nb%i == 0 || nb%(i+2) == 0 {
			return false
		}
		i += 6
	}
	return true
}
