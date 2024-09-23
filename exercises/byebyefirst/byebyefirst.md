# byebyefirst
## Instructions

Write a function that takes a slice of `string's` and returns a new slice without the first element.

  -  If the slice is empty, return the empty slice.

## Expected function

```
func ByeByeFirst(strings []string) []string {

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
	fmt.Println(ByeByeFirst([]string{}))
	fmt.Println(ByeByeFirst([]string{"one arg"}))
	fmt.Println(ByeByeFirst([]string{"first", "second"}))
	fmt.Println(ByeByeFirst([]string{"", "abcd", "efg"}))
}
```

And its output:

```
$ go run . | cat -e
[]$
[]$
[second]$
[abcd efg]$
```