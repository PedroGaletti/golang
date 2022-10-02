package functions

import "fmt"

// if you want export a function of your package it's necessary that the first letter be uppercase

func sum(x int, y int) int {
	return x + y
}

func returnTwoValues(x int, y int) (int, int) {
	return x + 1, y + 1
}

func quotient(x float64, y float64) (ans float64, err error) {
	if y == 0 {
		return 0, fmt.Errorf("you cant divide by zero")
	} else {
		return x / y, nil
	}
}

// Varadic functions
func sums(nums ...int) int {
	sum := 0

	for _, num := range nums {
		sum += num
	}

	return sum
}

func arraySum(arr []int) int {
	sum := 0

	for _, v := range arr {
		sum += v
	}

	return sum
}

func changeMyVar(myVar int) int {
	return myVar + 1
}

func changeMyVarWithPointer(myVar *int) int {
	return *myVar + 5
}

func Main() {
	// Simple function
	result := sum(1, 2)
	fmt.Println(result) // 3

	// Return more than one values in the function
	first, second := returnTwoValues(1, 2)
	fmt.Println(first, second) // 2, 3

	// Functions can return errors
	val, err := quotient(5, 4)
	fmt.Println("val", val) // 1.25
	fmt.Println("err", err) // nil

	fmt.Println(sums(1, 2, 3, 4, 5)) // 15

	// Array pass to function
	fmt.Println(arraySum([]int{1, 2, 3, 4, 5})) // 15

	// Functions don't change the original value of var outside like:
	myVar := 1
	fmt.Println("myVar before func changeMyVar: ", myVar) // 1
	fmt.Println(changeMyVar(myVar))                       // 2
	fmt.Println("myVar after func changeMyVar: ", myVar)  // 1

	// But with an pointer we can change the original value:
	fmt.Println("myVar before func changeVarWithPointer: ", myVar) // 1
	fmt.Println(changeMyVarWithPointer(&myVar))                    // 6
	fmt.Println("myVar after func changeVarWithPointer: ", myVar)  // 6
}
