package main

import (
	"fmt"
	"runtime"
	"strconv"
	"testing"
	"time"
)

/*
channels provide mechanism
- for communication
- for synchronisation
between goroutines

can cause goroutines to block and unblock
*/

/*
Properties
- goroutines can send and recieve data from channels.
- goroutine-safe
- FIFO mechanism
*/

// channel used for communication b/w goroutines
func TestBasicChannel(t *testing.T) {
	var ch chan string // declaration, nil channel by default
	ch = make(chan string)
	worker := func() {
		for i := 0; i < 5; i++ {
			work := <-ch // recieving from channel
			fmt.Printf("performing work, work - %v\n", work)
		}
	}
	manager := func() {
		for i := 0; i < 5; i++ {
			ch <- "clean house" + strconv.Itoa(i)
		}

		fmt.Println("done")
	}
	go worker()
	go manager()

}

/*
Unbuffered channel :
 - a sender will be blocked till a receiver is ready to receive from it.
*/
// iterate over channel
func TestRangeOverAChannel(t *testing.T) {
	var ch chan string // declaration, nil channel by default
	ch = make(chan string)
	worker := func() {
		for work := range ch {
			fmt.Printf("worker 0 performing work - %v\n", work)
		}
	}
	worker1 := func() {
		for work := range ch {
			fmt.Printf("worker 1 performing work - %v\n", work)
		}
	}
	manager := func() {
		for i := 0; i < 100; i++ {
			ch <- "clean house : " + strconv.Itoa(i)
		}
		//close(ch)
		fmt.Println("done")
	}
	go worker()
	go worker1()
	manager()
}

/*
If you dont close your channel, it will be GC when no longer required.
only necessary to close when a reciever is expecting it to close e.g the above case

*/

/*
if we try to read from a closed channel then it will read the default value
of channel type
*/
func TestClosedChannel(t *testing.T) {
	chaA := make(chan int)
	// chaB := make(chan string)

	go func() {
		for i := 0; i < 10; i++ {
			chaA <- i
		}
		close(chaA)
	}()

	for i := 0; i < 21; i++ {
		data := <-chaA

		fmt.Printf("data=%+v\n", data)
	}
}

/*
 select statement :

 select{
 case x := <-ch1 // do something with val x
 case y := <-ch2 // do something with val y
 default : it will execute when none of the channel is ready
 }
*/
func TestSelectStatement(t *testing.T) {

	chaA := make(chan int)
	chaB := make(chan int)
	quit := make(chan int)
	go func() {
		for i := 0; i < 10; i++ {
			chaA <- i
		}
		quit <- 0
		//close(chaA)
	}()
	go func() {
		for i := 10; i < 20; i++ {
			chaB <- i
		}
		//close(chaB)
	}()

	var done bool
	for {
		select {
		case x := <-chaA:
			fmt.Printf("received from %v\n", x)
			//time.Sleep(100 * time.Millisecond)
		case y := <-chaB:
			fmt.Printf("received from %v\n", y)
		case <-quit:
			done = true
		}
		if done {
			break
		}
	}

}

func TestNonMemoryLeakage(t *testing.T) {
	for i := 0; i < 5; i++ {
		TestFixMemoryLeakageBufferedChannel(t)
	}
	time.Sleep(time.Second)
	fmt.Println("no of goroutines in buffered channel", runtime.NumGoroutine()) // 7
}

// memory leakage
func TestMemoryLeakage(t *testing.T) {
	for i := 0; i < 10; i++ {
		TestMemoryLeakageNonBufferedChannel(t)
	}
	time.Sleep(time.Second)
	fmt.Println("no of goroutines in non buffered channel", runtime.NumGoroutine()) // 7

}

// fix memory Lekage using buffered channels

func TestFixMemoryLeakageBufferedChannel(t *testing.T) {
	chanA := make(chan int, 1)
	go func() {
		chanA <- 1
	}() // no one is consuming from it since it is buffered channel of capacity 1
	// Hence go routine will send to channel and will exit the goroutine -
	// not a blocking call
	done := make(chan int, 1)
	go func() {
		done <- 1
	}()
	time.Sleep(time.Second)
	select {
	case x := <-done:
		fmt.Println("x ", x)
	default:
		fmt.Println("default")
	}

}

// memory Leakage using non buffered channels
func TestMemoryLeakageNonBufferedChannel(t *testing.T) {
	chanA := make(chan int)
	go func() {
		chanA <- 1
	}() // no one is consuming from it but this func contains
	// pointer to this channel and hence will not be GC. - Memory Leakage
	done := make(chan int)
	go func() {
		done <- 1
	}()
	time.Sleep(time.Second)
	select {
	case x := <-done:
		fmt.Println("x ", x)
	default:
		fmt.Println("default")
	}

}

// Deadlock examples

func TestEmptySelect(t *testing.T) {
	select {}
	time.Sleep(10 * time.Second)
}

// Deadlock examples

func TestChannelJustReceiving(t *testing.T) {
	chanA := make(chan int, 1)

	<-chanA
}

//DataRace condition

func TestDataRace(t *testing.T) {
	work := func(i int) {
		fmt.Printf("i=%d\n", i)
	}
	for i := 0; i < 5; i++ {
		go func() {

			work(i) // data race, it will print whatever will be the value of i
			// at the point this func is called
		}()
	}
	time.Sleep(time.Second)
}

func TestFixDataRace(t *testing.T) {
	work := func(i int) {
		fmt.Printf("i=%d\n", i)
	}
	for i := 0; i < 5; i++ {
		j := i
		go func() {

			work(j)
		}()
	}
	time.Sleep(time.Second)
}
