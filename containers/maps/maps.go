package main

import "fmt"

func main() {
	// this is hash map
	m1 := map[string]string{
		"name": "Yubo",
		"Career": "SDE",
		"Age": "23",
	}

	m2 := make(map[string]int)
	fmt.Println(m1)
	fmt.Println(m2)

	fmt.Println("Traversing m1...")
	for k, v := range m1{
		fmt.Println(k, v)
	}

}
