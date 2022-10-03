package main

import "fmt"

/*
	Conditional operators: > < >= <= == !=
	Logical operators: && || !
*/

func IfConditionals() {
	fmt.Println("\nIfConditionals func: ")

	iAge := 8
	if (iAge >= 1) && (iAge <= 18) {
		fmt.Println("Important Birthday between 1 and 18")
	} else if (iAge == 21) || (iAge == 50) {
		fmt.Println("Important Birthday equals 21 or 50")
	} else if iAge >= 65 {
		fmt.Println("Important Birthday greater or equal 65")
	} else {
		fmt.Println("Not an important birthday")
	}

	fmt.Println("!true = ", !true)
}

func main() {
	IfConditionals()
}
