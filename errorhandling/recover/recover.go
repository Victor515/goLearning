package main

import (
	"fmt"
)

func tryRecover(){
	defer func(){
		r:= recover()
		if error, ok := r.(error); ok{
			fmt.Printf("An error has occured: %s", error.Error())
		} else{
			panic(r)
		}
	}()
	//panic(errors.New("This is an error!"))
	panic("Haha")
	//b := 0
	//a := 5 / b
	//fmt.Println(a)
}

func main() {
	tryRecover()
}
