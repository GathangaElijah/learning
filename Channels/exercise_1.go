package main

import (
	"fmt"
	"sync"
)

// Create a function that launches three goroutines that communicate using
// a channel. The first two goroutines each write 10 numbers to the
// channel. The third goroutine reads all the numbers from the channel and
// prints them out. The function should exit when all values have been
// printed out. Make sure that none of the goroutines leak. You can create
// additional goroutines if needed.

func Process() {
	ch := make(chan int)

	// Create 2 waitgroups
	// To signal when writing of data is over
	var wg sync.WaitGroup
	wg.Add(2)
	go func() {
		defer wg.Done()
		for i := 0; i < 10; i++ {
			ch <- i
		}
	}()

	go func() {
		defer wg.Done()
		for i := 0; i < 10; i++ {
			ch <- i*100 + 1
		}
	}()

	go func() {
		wg.Wait()
		close(ch)
	}()

	// the other waitgroup is to signal when reading of data is done.
	var wg2 sync.WaitGroup
	wg2.Add(1)
	go func() {
		defer wg2.Done()
		for v := range ch {
			fmt.Println(v)
		}
	}()
	wg2.Wait()
}

func main() {
	// Process()
	Launcher()
}
