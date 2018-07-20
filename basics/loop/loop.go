package main

import (
	"os"
	"bufio"
	"fmt"
)

func main() {
	const filename = "abc.txt"
	file, err := os.Open(filename)
	if err != nil{
		panic(err)
	}else{
		// create a scanner
		scanner := bufio.NewScanner(file)

		for scanner.Scan(){
			fmt.Println(scanner.Text())
		}
	}
}
