# weareunique
## Instructions

Write a function that takes two strings's and returns the number of characters that are not included in both, without repeating characters.

   - If there is no unique characters return `0`.
   - If both strings are empty return `-1`.

## Expected function

```
func WeAreUnique(str1 , str2 string) int {

}
```

## Usage

Here is a possible program to test your function:

```
package main

import (
	"fmt"
)

func main() {
	fmt.Println(WeAreUnique("foo", "boo"))
	fmt.Println(WeAreUnique("", ""))
	fmt.Println(WeAreUnique("abc", "def"))
}
```

And its output:

```
$ go run .
2
-1
6
```