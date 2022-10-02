package pointers

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

func ArrayPointers() {

}
