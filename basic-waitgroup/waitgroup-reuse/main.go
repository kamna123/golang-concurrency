package main

import "sync"

// here it will give error :
// panic: sync: WaitGroup is reused before previous Wait has returned
// bcz we are calling wg.Add before wg.Wait has finished the operation.
func main() {
	var wg sync.WaitGroup
	for i := 0; i < 100; i++ {
		go func() {
			wg.Add(3)
			go func() {
				defer wg.Done()
			}()
			go func() {
				defer wg.Done()
			}()
			go func() {
				defer wg.Done()
			}()
			wg.Wait() // here it will take time to update the count
			//	since it is doing some lengthy operations and we started adding tasks in WG
		}()
	}
}
