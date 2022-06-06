package main

import "fmt"

func main() {
	// Create a slice
	//var sl []int
	//sl := []int{0, 1, 2, 3, 4}
	//sl := make([]int, 5)
	//fmt.Printf("Slice: %v\n", sl)

	//////////////////////////////

	// Get value by index
	//index := 0
	//val := sl[index]
	//fmt.Printf("Value by index [%v]: %v\n", index, val)
	//val = sl[5]												//  -> panic: index out of range [5] with length 5

	//////////////////////////////

	// Get part of the slice
	//newOne := sl[1:]
	//fmt.Printf("newOne: %v\n", newOne)
	//fmt.Printf("sl: %v\n", sl)

	//////////////////////////////

	// Update value by index
	//sl[index] = 10
	//val := sl[index]
	//fmt.Printf("Updated value by index [%v]: %v\n", index, val)
	//sl[5] = 10												//  -> panic: index out of range [5] with length 5

	//////////////////////////////

	// Update v2
	//newOne := sl[1:]
	//fmt.Printf("newOne: %v\n", newOne)
	//fmt.Printf("sl: %v\n", sl)
	//
	//newOne[1] = -23
	//fmt.Printf("newOne: %v\n", newOne)
	//fmt.Printf("sl: %v\n", sl)

	//////////////////////////////

	// Go through all elements
	//for i, val := range sl {
	//	fmt.Printf("Index: %v - value: %v\n", i, val)
	//}

	//////////////////////////////

	// Append a new element
	//fmt.Printf("Before append: %v\n", sl)
	//sl = append(sl, -1)
	//fmt.Printf("After append: %v\n", sl)

	//////////////////////////////

	// Get length of the slice
	//length := len(sl)
	//fmt.Printf("Length of the slice: %v\n", length)

	//////////////////////////////

	// Get capacity of the slice
	//capacity := cap(sl)
	//fmt.Printf("Capacity of the slice: %v\n", capacity)

	//////////////////////////////

	// How does it grow
	//length := len(sl)
	//capacity := cap(sl)
	//fmt.Printf("Before append: %v\n", sl)
	//fmt.Printf("Length of the slice: %v\n", length)
	//fmt.Printf("Capacity of the slice: %v\n", capacity)

	//sl = append(sl, -1)

	//length = len(sl)
	//capacity = cap(sl)

	//fmt.Printf("After append: %v\n", sl)
	//fmt.Printf("Length of the slice: %v\n", length)
	//fmt.Printf("Capacity of the slice: %v\n", capacity)

	//////////////////////////////

	// Make a copy of slice
	//newOne := make([]int, 2)
	//count := copy(newOne, sl)
	//fmt.Printf("sl: %v\n", sl)
	//fmt.Printf("newOne: %v\n", newOne)
	//fmt.Printf("count: %v\n", count)

	//////////////////////////////

	// Using in functions
	//sl = make([]int, 5)
	//fmt.Printf("Slice before: %v\n", sl)
	//update1(sl)
	//fmt.Printf("Slice after: %v\n", sl)

	//sl = make([]int, 5)
	//fmt.Printf("Slice before: %v\n", sl)
	//update2(sl)
	//fmt.Printf("Slice after: %v\n", sl)
}

func update1(arr []int) {
	fmt.Printf("Slice in function before updating: %v\n", arr)
	arr[3] = -5
	fmt.Printf("Slice in function after updating: %v\n", arr)
}

func update2(arr []int) {
	fmt.Printf("Slice in function before updating: %v\n", arr)
	arr = append(arr, 10)
	fmt.Printf("Slice in function after updating: %v\n", arr)
}
