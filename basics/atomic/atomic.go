package main

import (
	"fmt"
	"time"
	"sync"
)

type atomicInt struct {
	value int
	lock sync.Mutex
}

func (a *atomicInt) increment(){
	fmt.Println("Block: safe increment")
	func(){
		a.lock.Lock()
		defer a.lock.Unlock() // good practice
		a.value++
	}()
}

func (a *atomicInt) get() int{
	a.lock.Lock()
	defer a.lock.Unlock()
	return a.value
}

func main() {
	var v atomicInt
	v.increment()
	go func() {
		v.increment()
	}()
	time.Sleep(time.Millisecond)
	fmt.Println("value: ", v.get())
}
