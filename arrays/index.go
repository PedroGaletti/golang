package arrays

import "fmt"

func Arrays() {
	var arr [5]int // Init the array with five empty spaces
	arr[0] = 1

	arr2 := [5]int{1, 2, 3, 4, 5} // Init the array with five spaces
	fmt.Println("arr2 index 0:", arr2[0])
	fmt.Println("arr length:", len(arr))

	for i := 0; i < len(arr2); i++ { // For loop in length of array
		fmt.Println(arr2[i])
	}

	for i, v := range arr2 { // For loop range in the array
		fmt.Println("%d : %d\n", i, v)
	}

	arr3 := [2][2]int{
		{1, 2},
		{3, 4},
	}

	for i := 0; i < 2; i++ {
		for x := 0; x < 2; x++ {
			fmt.Println(arr3[i][x])
		}
	}
}
