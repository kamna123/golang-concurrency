package main

import (
	"fmt"
	"time"
)

// sync exmaple taking elapsed : 406.435371ms
// To make is async we had to add 1 second of sleep, which is not
// an ideal solution. To handle it, we need to learn fork, join model
// main routine has forked a lot of child goroutine, so we need to join
// child routines to main. Refer fork-join module

func main() {
	now := time.Now()
	done := make(chan struct{})
	go task1(done) // fork
	go task2(done) // fork
	go task3(done) // fork
	go task4(done) // fork
	// send and recieved are blocking calls on the channel
	for i := 0; i < 4; i++ {
		<-done // join point
	}
	fmt.Println("elapsed :", time.Since(now))
}
func task1(done chan struct{}) {
	//time.Sleep(100 * time.Millisecond)
	fmt.Println("task1:")
	//<-done
	done <- struct{}{}

}
func task2(done chan struct{}) {
	//time.Sleep(200 * time.Millisecond)
	fmt.Println("task2:")
	done <- struct{}{}
}
func task3(done chan struct{}) {
	fmt.Println("task3:")
	done <- struct{}{}
}
func task4(done chan struct{}) {
	//	time.Sleep(100 * time.Millisecond)
	fmt.Println("task4:")
	done <- struct{}{}
}
