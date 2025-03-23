package main

import (
	"fmt"
	"sync"
)

// find&fix bug with reads-writes
func main() {
	storage := make(map[int]int, 1000)

	wg := sync.WaitGroup{}
	reads := 1000
	writes := 1000
	mu := sync.RWMutex{} // we need rw mutex to read lock

	wg.Add(writes)
	for i := 0; i < writes; i++ {
		i := i
		go func() {
			defer wg.Done()

			mu.Lock()
			defer mu.Unlock()
			storage[i] = i
		}()
	}
	wg.Add(reads)
	for i := 0; i < reads; i++ {
		i := i
		go func() {
			defer wg.Done()
			mu.RLock() // here is a usage of rlock and runlock. its more effective than usual mutex

			_, _ = storage[i]
			mu.RUnlock()
		}()
	}

	wg.Wait()
	fmt.Println(storage)
}
