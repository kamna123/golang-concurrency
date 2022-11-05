package main

import (
	"fmt"
	"sync"
	"testing"
)

// forgot to call wg.Done
// it will keep on waiting forever
// fatal error: all goroutines are asleep - deadlock!

func TestNotAddedDone(t *testing.T) {
	var wg sync.WaitGroup
	wg.Add(2)
	for i := 0; i < 2; i++ {
		j := i
		go func() {
			//defer wg.Done() // forgot to call it
			fmt.Printf("i = %d\n", j)
		}()
	}
	wg.Wait()

}

// fatal error: all goroutines are asleep - deadlock!
// you can add -race in command line to check race conditions.
func TestDoneIsLessThanAdd(t *testing.T) {
	var wg sync.WaitGroup
	wg.Add(2)
	for i := 0; i < 1; i++ {
		j := i
		go func() {
			defer wg.Done() //calling done once
			fmt.Printf("i = %d\n", j)
		}()
	}
	wg.Wait() // waiting for 2 tasks

}

// it will give error
// panic: sync: negative WaitGroup counter

func TestDoneIsMoreThanAdd(t *testing.T) {
	var wg sync.WaitGroup
	wg.Add(2)
	for i := 0; i < 3; i++ {
		j := i
		go func() {
			defer wg.Done() //calling done once
			fmt.Printf("i = %d\n", j)
		}()
	}
	wg.Wait() // waiting for 2 tasks

}
