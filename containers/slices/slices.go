package main

import "fmt"

func updateArr(arr []int){
	arr[0] = 100
}

func main() {

	// slice is nothing but a view of the original array
	arr := [...]int{0, 1, 2, 3, 4, 5}
	fmt.Println("arr = ", arr)

	// update arr using slice
	updateArr(arr[:])
	fmt.Println("updated arr = ", arr)

	// reslice
	s2 := arr[:]
	s2 = s2[:4]
	fmt.Println(s2)
	s2 = s2[1:]
	fmt.Println(s2)

	// extending slice
	// how slice is implemented? zero ptr, len, cap
	s1 := arr[2:4]
	s3 := s1[1:]
	fmt.Printf("s1 = %v, len(s1) = %d, cap(s1) = %d\n", s1, len(s1), cap(s1))
	fmt.Println("s3 = ", s3)

	// append to slice
	// if exceeds the cap, the system will create a new array for us
}
