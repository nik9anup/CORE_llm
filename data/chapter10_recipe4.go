// This program demonstrates the usage of sync.Pool to manage a pool of Worker objects,
// allowing efficient reuse of objects across multiple goroutines with minimized memory allocations.
package main

import (
	"fmt"
	"sync"
	"time"
)

type Worker struct {
	id string
}

func (w *Worker) String() string {
	return w.id
}

var globalCounter = 0

var pool = sync.Pool{
	New: func() interface{} {
		res := &Worker{fmt.Sprintf("%d", globalCounter)}
		globalCounter++
		return res
	},
}

func main() {
	wg := sync.WaitGroup{}
	wg.Add(10)
	for i := 0; i < 10; i++ {
		go func(idx int) {
			w := pool.Get().(*Worker) // Get a worker from the pool
			fmt.Println("Got worker ID: " + w.String())
			time.Sleep(time.Second) // Simulate work with the worker
			pool.Put(w)             // Put the worker back into the pool
			wg.Done()
		}(i)
	}
	wg.Wait()
}