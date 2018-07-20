package main

import "fmt"

func printSlice(s []int){
	fmt.Printf("slice = %v, len(s) = %d, cap(s) = %d\n", s, len(s), cap(s))
}

func main() {
	// create a slice
	var s []int
	for i := 0; i < 100; i++{
		// uncomment this line to see how slice is resized
		//printSlice(s)
		s = append(s, i)
	}

	s1 := []int{1, 3 ,4 ,6}
	printSlice(s1)

	s2 := make([]int, 16)
	s3 := make([]int, 10, 32)  // initialize with capacity
	printSlice(s2)
	printSlice(s3)

	// copy a slice
	fmt.Println("Copying s1 to s2")
	copy(s2, s1)
	printSlice(s2)

	// delete elements
	fmt.Println("Delete the 3rd elements from s2")
	s2 = append(s2[:3], s2[4:]...)
	printSlice(s2)
	fmt.Println("Popping at front")
	front := s2[0]
	s2 = s2[1: ]

	fmt.Println("Removing at end")
	tail := s2[len(s2) - 1]
	s2 = s2[:len(s2) - 1]
	fmt.Println(front, tail)
	printSlice(s2)

}
