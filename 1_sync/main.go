package main

import (
	"fmt"
	"sync"
)

// print square of range 0...20 in random order
func main() {
	counter := 20
	wg := sync.WaitGroup{}
	for i := 0; i < counter; i++ {
		wg.Add(1)
		go func(in int) {
			defer wg.Done()
			fmt.Println(i * i)
		}(i)
	}

	wg.Wait()
}
