// This code demonstrates the usage of sync.WaitGroup to synchronize and wait for a group of goroutines to complete,
// each printing an exit message with its index before signaling completion.
package main

import (
	"fmt"
	"sync"
)

func main() {
	wg := sync.WaitGroup{}
	for i := 0; i < 10; i++ {
		wg.Add(1)
		go func(idx int) {
			// Do some work
			defer wg.Done()
			fmt.Printf("Exiting %d\n", idx)
		}(i)
	}
	wg.Wait()
	fmt.Println("All done.")
}