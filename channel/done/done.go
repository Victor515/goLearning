package main

import (
	"fmt"
	"sync"
)

func doWork(id int, c chan int, done chan bool){
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
		done <- true // tell sender "I am done!"
	}
}

type worker struct {
	in chan int
	done chan bool
}

// channel as return value
func createWorker(id int) worker{
	w := worker{
		in : make(chan int),
		done: make(chan bool),
	}
	go doWork(id, w.in, w.done)
	return w
}


// Applying go built-in WaitGroup
func doWork2(id int, c chan int, wg *sync.WaitGroup){
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
		wg.Done()
	}
}

type worker2 struct {
	in chan int
	wg *sync.WaitGroup
}

func createWorker2(id int, wg *sync.WaitGroup) worker2{
	w := worker2{
		in : make(chan int),
		wg : wg,
	}
	go doWork2(id, w.in, w.wg)
	return w
}

func chanDemo(){
	//var c chan int // c == nil
	var workers [10]worker2

	// apply WaitGroup
	var wg sync.WaitGroup

	for i := 0; i < 10; i++{
		// create a channel for every worker and distribute to it
		//channels[i] = make(chan int)
		//go worker(i, channels[i])
		workers[i] = createWorker2(i, &wg) // pass the pointer here
	}

	wg.Add(20) // we have 20 tasks
	for i := 0; i < 10; i++{
		workers[i].in <- 'a' + i
		//<-workers[i].done // communicate by receiving from a channel
	}

	// wait for all
	/**
	Before sending the other information below, we need to receive "done" here
	This is because after sending information through channel, the routine will be blocked
	If we do not receive on the other end, a deadlock will be formed
	 */
	//for _, worker := range workers{
	//	<-worker.done
	//}

	for i := 0; i < 10; i++{
		workers[i].in <- 'A' + i
		//<-workers[i].done // communicate by receiving from a channel
	}

	//for _, worker := range workers{
	//	<-worker.done
	//	//<-worker.done
	//}
	wg.Wait()
}

func main() {
	// channel as first-class citizen
	chanDemo()
}


