package main

import (
	"fmt"
	"time"
)

func main() {
	now := time.Now()
	done := make(chan struct{})
	go func() { // fork point
		work()
		done <- struct{}{}
	}()
	<-done // join point
	fmt.Println("elapsed: ", time.Since(now))
	fmt.Println("done waiting, main exists")

}

func work() {
	time.Sleep(500 * time.Millisecond)
	fmt.Println("prinitng some stuff")
}
