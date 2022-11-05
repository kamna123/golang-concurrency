package main

import (
	"fmt"
	"runtime"
	"sync"
	"time"
)

func main() {

	fmt.Println("number of cores", runtime.NumCPU())
	var wg sync.WaitGroup
	wg.Add(10)
	now := time.Now()
	for i := 0; i < 10; i++ {
		go work(&wg, i)
	}
	wg.Wait() // will wait for all the tasks to get finished
	fmt.Println("elapsed time ", time.Since(now))
	fmt.Println("main is done")
}

func work(wg *sync.WaitGroup, id int) {
	defer wg.Done()
	//time.Sleep(100 * time.Millisecond)
	fmt.Println("task :", id, "is done")
}
