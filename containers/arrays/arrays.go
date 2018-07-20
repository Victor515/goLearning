package main

import "fmt"

func printArray(arr [4]int){
	arr[0] = 100
	for _, v := range arr{
		fmt.Println(v)
	}
}

func main() {
	var arr1 [5] int // declaration
	arr2 := [3]int{1, 3, 5}
	arr3 := [...]int{4, 5, 7, 10}

	var grid [4][5]int
	fmt.Println(arr1, arr2, arr3, grid)

	// iterate: traditional
	for i:= 0; i < len(arr2); i++{
		fmt.Println(arr2[i])
	}

	// iterate: range
	for i:= range arr2{
		fmt.Println(arr2[i])
	}

	for i, v:= range  arr2{
		fmt.Println(i, v)
	}

	for _, v:= range  arr2{
		fmt.Println(v)
	}

	// go array is pass by value(it will copy the whole array to the function namespace)
	fmt.Println("invoke printArray(arr3)")
	printArray(arr3);

	fmt.Println("arr3")
	for _, v := range arr3{
		fmt.Println(v)
	}

}
