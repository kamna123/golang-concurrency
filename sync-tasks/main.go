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
	go task1()
	go task2()
	go task3()
	go task4()
	time.Sleep(time.Second)
	// without join fork model
	fmt.Println("elapsed :", time.Since(now))
}
func task1() {
	time.Sleep(100 * time.Millisecond)
	fmt.Println("task1:")
}
func task2() {
	time.Sleep(200 * time.Millisecond)
	fmt.Println("task2:")
}
func task3() {
	fmt.Println("task3:")
}
func task4() {
	time.Sleep(100 * time.Millisecond)
	fmt.Println("task4:")
}
