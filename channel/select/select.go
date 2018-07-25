package main

import (
	"time"
	"math/rand"
	"fmt"
)

/**
Use select to coordinate different channels
time.After help us do something after a given period
time.Tick help us do something for every given period

 */

 /**
 Below is a reader/writer example, for every random(0,1500) millisecond, two channels will both send a value,
 and for every second, the worker will receive one. Because the producing rate is faster than consuming rate, we
 need to queue to store not consumed data.
  */

func generator() chan int{
	out := make(chan int)
	go func() {
		i := 0
		for{
			time.Sleep(time.Duration(rand.Intn(1500)) * time.Millisecond)
			out <- i
			i++
		}
	}()

	return out
}

func worker(id int, c chan int){
	for n := range c{
		time.Sleep(time.Second)
		fmt.Printf("Worker %d receives %d\n", id, n)
	}
}

// channel as return value
func createWorker(id int) chan<- int{
	c := make(chan int)
	go worker(id, c) // create a worker
	return c
}

func main() {
	var c1, c2 chan int = generator(), generator()
	w := createWorker(0)
	// non-block sending/receiving value

	var values []int

	// create a timer
	tm := time.After(10 * time.Second)
	tick := time.Tick(time.Second)
	for {
		var activeWorker chan <-int // this is nil channel! select will not select a nil channel
		var activeValue int
		if len(values) > 0{
			activeWorker = w
			activeValue = values[0]
		}
		select {
		case n := <-c1:
			values = append(values, n)
		case n := <-c2:
			values = append(values, n)
		case activeWorker <- activeValue:
			values = values[1:]
		// if no data is received within 800 millisecond
		case <-time.After(800 * time.Millisecond):
			fmt.Println("Timeout")
		case <- tick:
			fmt.Println("queue length is: ", len(values))
		case <-tm:
			fmt.Println("Bye!")
			return
		//default:
		//	fmt.Println("No value received")
		}
	}
}
