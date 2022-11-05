package main

import (
	"fmt"
	"time"
)

// it will not print 'prinitng some stuff', bcz join point is missing
func main() {
	go work() // fork point
	time.Sleep(100 * time.Millisecond)
	fmt.Println("done waiting, main exists")
	// join point
}

func work() {
	time.Sleep(500 * time.Millisecond)
	fmt.Println("prinitng some stuff")
}
