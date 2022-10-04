package main

import "fmt"

type Constraint interface {
	int | float64
}

func SumWithConstraint[T Constraint](x T, y T) T {
	return x + y
}

func main() {
	fmt.Println("\nConstraints func: ")
	fmt.Println("5 + 4 = ", SumWithConstraint(5, 4))         // 5 + 4 = 9
	fmt.Println("5.6 + 4.7 = ", SumWithConstraint(5.6, 4.7)) // 5.6 + 4.7 = 10.3
}
