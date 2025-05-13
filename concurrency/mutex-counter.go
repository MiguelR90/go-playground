package main

import (
	"fmt"
	"sync"
	"time"
)

// SafeCounter is safe to use concurrently.
type SafeCounter struct {
	mu sync.Mutex
	v  map[string]int
}

// Inc increments the counter for the given key.
func (c *SafeCounter) Inc(key string) {
	// Defer can be used just like python with / context managers
	// it helps you define blocks with setup and teardown activities into a single context
	c.mu.Lock()
	defer c.mu.Unlock()

	// Lock so only on goroutine at a time can access the map c.v
	c.v[key]++
}

func (c *SafeCounter) Value(key string) int {
	// defer as resource manager / cleanup
	c.mu.Lock()
	defer c.mu.Unlock()

	return c.v[key]
}

func main() {
	c := SafeCounter{v: make(map[string]int)}

	// this kicks off 1000 goroutines that will all try to increment the counter
	for i := 0; i < 1000; i++ {
		go c.Inc("somekey")
	}

	// i guess we need the sleep to make sure that all the go routines have completed?
	time.Sleep(time.Second)
	fmt.Println(c.Value("somekey"))
}
