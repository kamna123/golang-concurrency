package main

import (
	"fmt"
	"sync"
	"sync/atomic"
	"testing"
)

// counter value is not 1000 as expected
// bcz critical section is not synchronized and is shared by threads
func TestConcurrentWrites(t *testing.T) {
	counter := 0
	var wg sync.WaitGroup
	for i := 0; i < 1000; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			counter++
		}()
	}
	wg.Wait()
	fmt.Printf("Counter- %+v\n", counter)
}

// solution

// use sync.atomic

func TestConcurrentWritesUsingSyncAtomic(t *testing.T) {
	counter := int32(0)
	var wg sync.WaitGroup
	for i := 0; i < 1000; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			atomic.AddInt32(&counter, 1) // this is atomic operation
		}()
	}
	wg.Wait()

	fmt.Printf("Counter - %+v\n", counter)
}

// solution 2
// use mutex
func TestConcurrentWritesUsingMutex(t *testing.T) {
	counter := int32(0)
	var wg sync.WaitGroup
	mu := sync.Mutex{}
	for i := 0; i < 1000; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			mu.Lock()
			counter++
			mu.Unlock()
		}()
	}
	wg.Wait()

	fmt.Printf("Counter - %+v\n", counter)
}

// concurrent map
// it will give error fatal error: concurrent map writes

func TestConcurrentWriteMapBuggy(t *testing.T) {
	m := map[int]int{}
	var wg sync.WaitGroup
	for i := 0; i < 1000; i++ {
		wg.Add(1)
		j := i
		go func() {
			defer wg.Done()
			m[j] = 1
		}()
	}
	for i := 0; i < 1000; i++ {
		wg.Add(1)
		j := i
		go func() {
			defer wg.Done()
			if _, ok := m[j]; ok {

			}

		}()
	}
	wg.Wait()

}

// solution 1 - use mutex

func TestConcurrentWriteMapFixed(t *testing.T) {
	m := map[int]int{}
	mu := sync.RWMutex{}
	var wg sync.WaitGroup
	for i := 0; i < 1000; i++ {
		wg.Add(1)
		j := i
		go func() {
			defer wg.Done()
			mu.Lock()
			m[j] = 1
			mu.Unlock()
		}()
	}
	for i := 0; i < 1000; i++ {
		wg.Add(1)
		j := i
		go func() {
			defer wg.Done()
			mu.RLock() // use read lock for reading
			if _, ok := m[j]; ok {

			}
			mu.RUnlock()

		}()
	}
	wg.Wait()

}

// solution 2

// define read and write for map as separate func using mutex.lock
// for cleaner code

type ConcurrentMap struct {
	m  map[int]int
	mu sync.RWMutex
}

func NewConcurrentMap() *ConcurrentMap {
	return &ConcurrentMap{
		m:  map[int]int{},
		mu: sync.RWMutex{},
	}
}

func (c *ConcurrentMap) Put(k, v int) {
	c.mu.Lock()
	defer c.mu.Unlock()
	c.m[k] = v
}

func (c *ConcurrentMap) Get(k int) int {
	c.mu.RLock()
	defer c.mu.RUnlock()
	return c.m[k]
}

// solution 3 use sync.Map -> see implementation details
