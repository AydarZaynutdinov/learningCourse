package main

import "fmt"

func main() {
	// Create an array
	//var arr [5]int
	//arr := [5]int{}
	//arr := [5]int{0, 1, 2, 3, 4}
	//arr := [5]int{0, 1, 2}
	//arr := [5]int{0, 1, 2, 3, 4, 5} 							// -> error
	//fmt.Printf("Array: %v\n", arr)

	//////////////////////////////

	// Get value by index
	//index := 0
	//val := arr[index]
	//_ = arr[5]												// -> error
	//fmt.Printf("Value by index [%v]: %v\n", index, val)

	//////////////////////////////

	// Update value by index
	//arr[index] = 10
	//val := arr[index]
	//arr[5] = 10												// -> error
	//fmt.Printf("Updated value by index [%v]: %v\n", index, val)

	//////////////////////////////
	// Go through all elements
	//for i, val := range arr {
	//	fmt.Printf("Index: %v - value: %v\n", i, val)
	//}

	//////////////////////////////

	// Get length of an array
	//length := len(arr)
	//fmt.Printf("Length of the array: %v\n", length)

	//////////////////////////////

	// Using in functions
	//arr = [5]int{}
	//fmt.Printf("Array before: %v\n", arr)
	//update1(arr)
	//fmt.Printf("Array after: %v\n", arr)

	//arr = [5]int{}
	//fmt.Printf("Array before: %v\n", arr)
	//update2(&arr)
	//fmt.Printf("Array after: %v\n", arr)
}

func update1(arr [5]int) {
	fmt.Printf("Array in function before updating: %v\n", arr)
	arr[3] = -5
	fmt.Printf("Array in function after updating: %v\n", arr)
}

func update2(arr *[5]int) {
	fmt.Printf("Array in function before updating: %v\n", arr)
	arr[3] = -5
	fmt.Printf("Array in function after updating: %v\n", arr)
}
