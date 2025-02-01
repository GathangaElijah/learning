package main

import (
	"fmt"
	"sync"
)

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
	Process()
}
