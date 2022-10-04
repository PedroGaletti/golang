package main

import "fmt"

func useFunc(f func(int, int) int, x, y int) {
	fmt.Println("Answer: ", (f(x, y)))
}

func sum(x, y int) int {
	return x + y
}

func main() {
	intSum := func(x, y int) int { return x + y }
	fmt.Println("5 + 4 = ", intSum(5, 4)) // 5 + 4 = 9

	vInt := 1
	changeVar := func() {
		vInt += 1
	}
	changeVar()
	fmt.Println("vInt = ", vInt) // vInt = 2

	useFunc(sum, 5, 8) // Answer: 13
}
