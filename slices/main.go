package main

import "fmt"

func SlicesRunesAndStrings() {
	str := "abcde"
	arr := []rune(str) // Transforming characters of the string in the unicodes
	for _, v := range arr {
		fmt.Printf("Rune array: %d\n", v) // Loop the array of runes
	}
	bytes := []byte{'a', 'b', 'c'}
	bStr := string(bytes[:]) // Transforming the array of bytes into string
	fmt.Println("I'm a string: ", bStr)
}

func SlicesDataTypes() {
	// var name []datatype
	sl1 := make([]string, 6)
	sl1[0] = "Society"
	sl1[1] = "of"
	sl1[2] = "the"
	sl1[3] = "Simulated"
	sl1[4] = "Universe"
	fmt.Println("Slice size: ", len(sl1))

	for i := 0; i < len(sl1); i++ {
		fmt.Println(sl1[i])
	}

	for _, x := range sl1 {
		fmt.Println(x)
	}

	arr := [5]int{1, 2, 3, 4, 5}
	vSlice := arr[0:2]
	fmt.Println("1st 2: ", arr[0:2])                         // [1, 2]
	fmt.Println("1st 3: ", arr[:3])                          // [1, 2, 3]
	fmt.Println("Last 3: ", arr[2:])                         // [3, 4, 5]
	fmt.Println("Exclude last position: ", arr[:len(arr)-1]) // [1, 2, 3, 4]
	fmt.Println("Last position: ", arr[len(arr)-1])          // 5

	arr[0] = 10                     // When change the original array the slice var will change:
	fmt.Println("vSlice: ", vSlice) // [10, 2]
	vSlice[0] = 1                   // When change the slice var the original array will change:
	fmt.Println("arr: ", arr)       // [1, 2, 3, 4, 5]

	vSlice = append(vSlice, 6)      // append just in the slice array
	fmt.Println("vSlice: ", vSlice) // [1, 2, 6]
	fmt.Println("arr: ", arr)       // [1, 2, 6, 4, 5]

	vArrSlice := make([]string, 6)
	fmt.Println("vArrSlice", vArrSlice)    // [      ]
	fmt.Println("vArrSlice", vArrSlice[0]) // [ ]
}

func main() {
	SlicesRunesAndStrings()
	SlicesDataTypes()
}
