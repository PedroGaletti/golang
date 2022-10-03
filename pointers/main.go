package main

import "fmt"

func changeVal(val *int) int {
	return *val + 1
}

func Pointers() {
	val := 10
	var valPtr *int = &val
	fmt.Println("Val adress: ", valPtr) // 0xc000.....
	fmt.Println("Val value: ", *valPtr) // Print the value of pointer: 10

	*valPtr = 11                        // Change value of pointer
	fmt.Println("Val value: ", *valPtr) // Value of the pointer: 11

	fmt.Println("Val before func: ", val) // var value before the function 11
	changeVal(&val)                       // Turn the var into pointer and pass to the function
	fmt.Println("Val after func: ", val)  // Return the value of var val 12
}

func dblArrVals(arr *[4]int) {
	x := 0
	for x < 4 {
		arr[x] *= 2
		x++
	}
}

func average(nums ...float64) float64 {
	// It's not possible divide float64 / int
	// Need be float64 / float64

	var sum float64 = 0.0
	var size float64 = float64(len(nums)) // Transform size to float64

	for _, val := range nums {
		sum += val
	}

	return sum / size
}

func ArrayPointers() {
	arr := [4]int{1, 2, 3, 4}
	dblArrVals(&arr) // Pass the pointer array
	fmt.Println(arr) // [2, 4, 6, 8]

	// Passing a slice in an function
	slice := []float64{11, 13, 17}
	fmt.Printf("Average: %.3f\n", average(slice...))
}

func main() {
	Pointers()
	ArrayPointers()
}
