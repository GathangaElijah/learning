package main

import "fmt"

func isPrime(n int) bool {
	if n <= 1 {
		return false
	} else if n <= 3 {
		return true
	} else if n%2 == 0 || n%3 == 0 {
		return false
	}

	i := 5
	for i*i < n {
		if n%i == 0 || n%(i+2) == 0 {
			return false
		}
		i += 6
	}
	return true
}

func FindPrevPrime(nb int) int {
	var primes []int
	if nb <= 1 {
		return 0
	}
		for i := 1; i <= nb; i ++{
			if isPrime(i){
				primes = append(primes, i)
			}
		}
		if !isPrime(nb){
			return primes[len(primes)-1]
		}
		return nb
}

func main() {
	fmt.Println(FindPrevPrime(3000))
	fmt.Println(FindPrevPrime(1))
}
