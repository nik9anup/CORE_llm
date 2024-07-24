// This code defines a Source type with a Pop method that ensures data loading occurs only once,
// simulating a delayed initialization of data with sync.Mutex and sync.Once synchronization mechanisms,
// and demonstrates concurrent access to the Pop method by multiple goroutines.
package main

import (
	"fmt"
	"sync"
	"time"
)

var names = []interface{}{"Alan", "Joe", "Jack", "Ben",
	"Ellen", "Lisa", "Carl", "Steve", "Anton", "Yo"}

type Source struct {
	m    *sync.Mutex
	o    *sync.Once
	data []interface{}
}

func (s *Source) Pop() (interface{}, error) {
	s.m.Lock()
	defer s.m.Unlock()
	s.o.Do(func() {
		time.Sleep(time.Second * 30) // Simulates data loading delay
		s.data = names
		fmt.Println("Data has been loaded.")
	})
	if len(s.data) > 0 {
		res := s.data[0]
		s.data = s.data[1:]
		return res, nil
	}
	return nil, fmt.Errorf("No data available")
}

func main() {

	s := &Source{&sync.Mutex{}, &sync.Once{}, nil}
	wg := sync.WaitGroup{}
	wg.Add(10)
	for i := 0; i < 10; i++ {
		go func(idx int) {
			if val, err := s.Pop(); err == nil {
				fmt.Printf("Pop %d returned: %s\n", idx, val)
			}
			wg.Done()
		}(i)
	}
	wg.Wait()
}