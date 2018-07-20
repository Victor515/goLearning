package main

import (
	"fmt"
	"math/cmplx"
	"math"
)

func variableZeroValue(){
	var a int
	var b string
	c := 1000
	d := "yes"
	fmt.Printf("%d %q", a, b)
	fmt.Println(c ,d)
}

func euler(){
	fmt.Println(cmplx.Pow(math.E, 1i * math.Pi) + 1)
}

func enums(){
	const(
		cpp = iota
		java
		python
		golang
	)

	fmt.Println(cpp, java, python, golang)
}

func main() {
	fmt.Printf("learngo, world\n")
	variableZeroValue()
	euler()
	enums()
}