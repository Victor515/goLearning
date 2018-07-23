package main

import (
	"fmt"
	"io"
	"bufio"
	"strings"
	"learngo/functional/fib"
)


type intGen func() int

// build a interface for function() int as type
// GO语言中，函数是一等公民
func (g intGen) Read(p []byte) (n int, err error) {
	next := g()

	if(next > 10000){
		return 0, io.EOF
	}

	// convert to string
	s := fmt.Sprintf("%d\n", next)

	// TODO: Incorerct if p is too small

	// use string readers to read
	return strings.NewReader(s).Read(p)
}

// take a Reader and print its contents
func printFileContents(reader io.Reader){
	scanner := bufio.NewScanner(reader)

	for scanner.Scan(){
		fmt.Println(scanner.Text())
	}
}

func main() {
	var f intGen = fib.Fibonacci()

	//for i := 0; i < 10; i++{
	//	fmt.Println(f())
	//}

	printFileContents(f)
}
