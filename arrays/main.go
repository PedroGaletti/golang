package main

import "fmt"

func Arrays() {
	fmt.Println("\nArrays func: ")

	var arr [5]int // Init the array with five empty spaces
	arr[0] = 1

	arr2 := [5]int{1, 2, 3, 4, 5}         // Init the array with five spaces
	fmt.Println("arr2 index 0:", arr2[0]) // Print: 1
	fmt.Println("arr length:", len(arr))  // Print: 5

	for i := 0; i < len(arr2); i++ { // For loop in length of array
		fmt.Print(arr2[i]) // Print: 1 2 3 4 5
	}

	fmt.Printf("\n") // Just to separate the logs

	for i, v := range arr2 { // For loop range in the array
		fmt.Printf(" %d:%d |", i, v) // Print: 0 : 1 | 1 : 2 | 2 : 3 | 3 : 4 | 4 : 5 |
	}

	fmt.Printf("\n") // Just to separate the logs

	arr3 := [2][2]int{
		{1, 2},
		{3, 4},
	}

	for i := 0; i < 2; i++ {
		for x := 0; x < 2; x++ {
			fmt.Print(arr3[i][x]) // Print: 1 2 3 4
		}
	}

	fmt.Printf("\n") // Just to separate the logs
}

func main() {
	Arrays()
}
