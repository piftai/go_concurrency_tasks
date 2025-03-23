package main

import (
	"fmt"
	"sync"
)

// find and fix 2 bugs
func main() {
	storage := make(map[int]int) // first bug is here. it's nil map

	mu := sync.Mutex{} // second bug is here. we did not have mutex, but we need it
	wg := sync.WaitGroup{}
	writes := 1000

	wg.Add(writes)
	for i := 0; i < writes; i++ {
		i := i
		go func() {
			defer wg.Done()
			mu.Lock()
			storage[i] = i
			mu.Unlock()
		}()
	}

	wg.Wait()
	fmt.Println(storage)
}
