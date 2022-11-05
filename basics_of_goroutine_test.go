package main

import (
	"fmt"
	"testing"
	"time"
)

/*

Goroutine : lightweight threads managed by go-routine
 - user-space threads
 - Go schedular schedules these threads on OS threads at runtime (M-N mappings)
 - Typical use cases :
	- where we need to execute task concurrently
- Goroutines share same address space, so there should be mutual exclusion in critical sections.
*/
// run multiple goroutines
func TestBasicGoRoutine(t *testing.T) {
	sayHi := func() {
		fmt.Printf("Say Hiiiiii\n")
	}
	sayHello := func() {
		fmt.Printf("Say hello\n")
	}

	go sayHi()
	go sayHello()
	go sayHi()
	go sayHello()
	time.Sleep(time.Second) // it is required else main goroutine will exit before execution og above goroutines
}

/*
Notice few things
- when you call go f(a,b,c), the values of a, b, c are evaluated in the current
  go routine context
- and Func f runs in a separate thread
*/
