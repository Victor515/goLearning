package main

import (
	"time"
	"fmt"
)

func worker(id int, c chan int){
	//for{
	//	n, ok := <- c
	//	if !ok{
	//		break
	//	}
	//	fmt.Printf("Worker %d receives %c\n", id, n)
	//}

	// range is always ok
	for n := range c{
		fmt.Printf("Worker %d receives %c\n", id, n)
	}
}

// channel as return value
func createWorker(id int) chan<- int{
	c := make(chan int)
	go worker(id, c)
	return c
}

func chanDemo(){
	//var c chan int // c == nil
	var channels [10]chan<- int
	for i := 0; i < 10; i++{
		// create a channel for every worker and distribute to it
		//channels[i] = make(chan int)
		//go worker(i, channels[i])
		channels[i] = createWorker(i)
	}

	for i := 0; i < 10; i++{
		channels[i] <- 'a' + i
	}

	for i := 0; i < 10; i++{
		channels[i] <- 'A' + i
	}

	time.Sleep(time.Millisecond)
}


func bufferChannel(){
	c := make(chan int, 3) // create a buffer of size 3

	c <- 'a' // sending to channel won't immediately create a switch
	c <- 'b'
	c <- 'c'

	go worker(0, c)
	c <- 'd'
	time.Sleep(time.Millisecond)
}

func channelClose() {
	c := make(chan int, 3) // create a buffer of size 3

	c <- 'a' // sending to channel won't immediately create a switch
	c <- 'b'
	c <- 'c'

	go worker(0, c)
	c <- 'd'
	close(c)
	time.Sleep(time.Millisecond)
}

func main() {
	// channel as first-class citizen
	chanDemo()
	// buffered channel
	bufferChannel()
	// close channel
	channelClose()
}


