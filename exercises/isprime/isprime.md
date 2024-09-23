# isPrime
## Instructions
Write a function that returns `true` if the `int` passed as parameter is a `prime` number. Otherwise it returns `false`.<br>

The function must be optimized in order to avoid time-outs with the tester.<br>

(We consider that only `positive` numbers can be prime numbers)<br>

(We also consider that `1` is not a prime number)<br>
## Expected function
```
func IsPrime(nb int) bool {

}

```
## Usage
here is a possible program to test your function
```
package main

import (
	"fmt"
	"piscine"
)

func main() {
	fmt.Println(piscine.IsPrime(5))
	fmt.Println(piscine.IsPrime(4))
}

```
and its output
```
$ go run .
true
false
$

```