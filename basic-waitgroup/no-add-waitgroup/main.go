package main

import (
	"fmt"
	"sync"
	"time"
)

// without add func with return immediately
// since it is not waiting for any task

func main() {
	var wg sync.WaitGroup
	go func() {
		defer wg.Done()
		time.Sleep(100 * time.Millisecond)
		fmt.Println("In func")
	}()
	wg.Wait()
	//	time.Sleep(time.Second)
	fmt.Println("returned immediately")

}
