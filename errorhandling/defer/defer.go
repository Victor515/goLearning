package main

import (
	"os"
	"learngo/functional/fib"
	"fmt"
	"bufio"
)

func writeFile(filename string){
	file, err := os.Create(filename)
	if err != nil{
		panic(err)
	}

	// defer to close file
	defer file.Close()

	// create a buffer writer
	writer := bufio.NewWriter(file)
	// flush to disk
	defer writer.Flush()

	f := fib.Fibonacci()

	for i := 0; i < 20; i++ {
		fmt.Fprintln(writer, f())
	}
}

func main() {
	writeFile("fib.txt")
}
