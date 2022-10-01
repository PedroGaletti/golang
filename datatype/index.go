package datatype

import (
	"fmt"
	"reflect"
	"strconv"
)

func MainDataTypes() {
	/*
		Datatypes: int, float64, bool, string, rune
		Default values: 0, 0.0, false, ""
	*/
	fmt.Println(reflect.TypeOf((25)))      // Int
	fmt.Println(reflect.TypeOf((3.14)))    // Float64
	fmt.Println(reflect.TypeOf((true)))    // Bool
	fmt.Println(reflect.TypeOf(("Hello"))) // String
	fmt.Println(reflect.TypeOf(('🚀')))     // Int32
}

func ConvertDataTypes() {
	cv1 := 1.5                       // Float64
	cv2 := int(cv1)                  // Converting float64 into int
	fmt.Println(reflect.TypeOf(cv2)) // Print

	cv3 := "500000"
	cv4, err := strconv.Atoi(cv3)              // Converting string into int
	fmt.Println(cv4, err, reflect.TypeOf(cv4)) // Print

	cv5 := 500000
	cv6 := strconv.Itoa(cv5)              // Converting int into string
	fmt.Println(cv6, reflect.TypeOf(cv6)) // Print

	cv7 := "3.14"
	if cv8, err := strconv.ParseFloat(cv7, 64); err == nil { // Converting a string into float64
		fmt.Println(cv8, reflect.TypeOf(cv8)) // Print
	}

	cv9 := fmt.Sprintf("%f", 3.14)        // Converting a float64 into string
	fmt.Println(cv9, reflect.TypeOf(cv9)) // Print
}

func PrintDataTypes() {
	/*
		%d: Integer
		%c: Character
		%f: Float
		%t: Boolean
		%s: String
		%o: Base 8
		%x: Base 16
		%v: Guesses based on data type
		%T: Type of supplied value
	*/

	fmt.Printf("%s %d %c %f %t %o %x\n", "stuff", 1, 'a', 3.14, true, 1, 1)
	fmt.Printf("%9f\n", 3.14)              // Nine characters in the end
	fmt.Printf("%.2f\n", 3.141592)         // Two decimals
	fmt.Printf("%9.f\n", 3.141592)         // Nine space in the start
	sp1 := fmt.Sprintf("%9.f\n", 3.141592) // Nine space in the start
	fmt.Println(sp1)                       // Print
}
