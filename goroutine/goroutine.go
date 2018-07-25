package main

import (
	"fmt"
	"time"
)

// Goroutine is more like coroutine than thread
// use top command, we can check how many threads are running at the same time
func main() {
	for i := 0; i < 1000; i++{
		go func(i int){
			for {
				fmt.Printf("Hello from goroutine: %d\n", i)
			}
		}(i)
		//time.Sleep(time.Millisecond)
		time.Sleep(time.Minute)
	}
}

//func main(){
//	var a [10]int
//	for i:= 0; i < 10; i++{
//		go func(i int) {
//			for{
//				a[i]++
//				runtime.Gosched() //yield control
//			}
//		}(i)
//	}
//	time.Sleep(time.Millisecond)
//	fmt.Println(a)
//}

/**
Use -race to detect Data Race
 */
//func main(){
//	var a [10]int
//	for i:= 0; i < 10; i++{
//		go func() { // race condition
//			for{
//				a[i]++
//				runtime.Gosched() //yield control
//			}
//		}()
//	}
//	time.Sleep(time.Millisecond)
//	fmt.Println(a)
//}
