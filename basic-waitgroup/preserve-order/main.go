package main

import (
	"fmt"
	"sync"
)

// order can be 1, {2,3}, 4 // 2,3 can be in any order
func main() {
	var wg sync.WaitGroup
	wg.Add(1)
	go task1(&wg)
	wg.Wait()

	wg.Add(2)
	go task3(&wg)
	go task2(&wg)
	wg.Wait()

	wg.Add(1)
	go task4(&wg)
	wg.Wait()
}

func task1(wg *sync.WaitGroup) {
	defer wg.Done()
	fmt.Println("task 1")
}

func task2(wg *sync.WaitGroup) {
	defer wg.Done()
	fmt.Println("task 2")
}

func task3(wg *sync.WaitGroup) {
	defer wg.Done()
	fmt.Println("task 3")
}

func task4(wg *sync.WaitGroup) {
	defer wg.Done()
	fmt.Println("task 4")
}
