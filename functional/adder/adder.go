package main

import "fmt"

// 函数体中存在着局部变量和自由变量
func adder() func(int) int{
	sum := 0 // 自由变量

	// 返回的是函数，以及它的闭包
	return func(v int) int {
		sum += v // v 是局部变量
		return sum
	}
}

// following "functional" principles
type iAdder func(int) (int, iAdder)

func adder2(base int) iAdder {
	return func(v int) (int, iAdder){
		return base + v, adder2(base + v)
	}
}

func main() {
	a := adder()
	for i := 0; i < 10; i++{
		fmt.Printf("sum from 0 to %d is %d\n", i, a(i))
	}
}
