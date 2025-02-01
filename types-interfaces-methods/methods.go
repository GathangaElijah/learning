package main

import (
	"fmt"
	"time"
)

//Demonstrating the pointer and value receivers.

type Counter struct {
	total int
	lastUpdated time.Time
}

// This demonstrates a pointer receiver in a method.
func (c *Counter) Increment() {
	c.total++
	c.lastUpdated = time.Now()
}

// This demonstrates a value receiver in method.
func (c Counter) String() string {
	return fmt.Sprintf("total: %d, last updated: %v", c.total, c.lastUpdated)
}

func main(){
	// calling on value type
	var c Counter
	fmt.Println("Am calling this on a value type >>\n", c.String())
	c.Increment()
	fmt.Println("Am calling this on a value type >>\n", c.String())

	// calling on a pointer type
	var c1 = &Counter{}
	fmt.Println("Am calling this on a pointer type >>\n", c1.String())
	c1.Increment()
	fmt.Println("Am calling this on a pointer type >>\n", c1.String())

}