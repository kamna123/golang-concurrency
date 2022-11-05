package main

import (
	"fmt"
	"runtime"
	"time"
)

func main() {

	fmt.Println("number of cores", runtime.NumCPU())
	for i := 0; i < 10; i++ {
		go work(i)
	}
	time.Sleep(100 * time.Millisecond)
	fmt.Println("main is done")
}

func work(id int) {
	time.Sleep(100 * time.Millisecond)
	fmt.Println("task :", id, "is done")
}
