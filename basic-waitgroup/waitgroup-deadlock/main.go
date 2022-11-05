package main

import "sync"

// if we call done less than the time add contains value
// it will lead to deadlock bcz we are waiting for something
// which will not be happening.

//fatal error: all goroutines are asleep - deadlock!

func main() {
	var wg sync.WaitGroup
	wg.Add(2)
	wg.Wait()

}
