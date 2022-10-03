package main

import (
	"fmt"
	"reflect"
	"strconv"
)

/*
	Datatypes: int, float64, bool, string, rune
	Default values: 0, 0.0, false, ""

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

func MainDataTypes() {
	fmt.Println("\nMainDataTypes func: ")

	fmt.Println(reflect.TypeOf((25)))      // Print: Int
	fmt.Println(reflect.TypeOf((3.14)))    // Print: Float64
	fmt.Println(reflect.TypeOf((true)))    // Print: Bool
	fmt.Println(reflect.TypeOf(("Hello"))) // Print: String
	fmt.Println(reflect.TypeOf(('🚀')))     // Print: Int32
}

func ConvertDataTypes() {
	fmt.Println("\nConvertDataTypes func: ")

	cv1 := 1.5                       // Float64
	cv2 := int(cv1)                  // Converting float64 into int
	fmt.Println(reflect.TypeOf(cv2)) // Print: Int

	cv3 := "500000"
	cv4, err := strconv.Atoi(cv3)              // Converting string into int
	fmt.Println(cv4, err, reflect.TypeOf(cv4)) // Print: 500000 <nil> int

	cv5 := 500000
	cv6 := strconv.Itoa(cv5)              // Converting int into string
	fmt.Println(cv6, reflect.TypeOf(cv6)) // Print: 500000 string

	cv7 := "3.14"
	if cv8, err := strconv.ParseFloat(cv7, 64); err == nil { // Converting a string into float64
		fmt.Println(cv8, reflect.TypeOf(cv8)) // Print: 3.14 float64
	}

	cv9 := fmt.Sprintf("%f", 3.14)        // Converting a float64 into string
	fmt.Println(cv9, reflect.TypeOf(cv9)) // Print: 3.140000 string
}

func PrintDataTypes() {
	fmt.Println("\nPrintDataTypes func: ")
	fmt.Printf("%s %d %c %.2f %t %o %x\n", "Stuff", 1, 'a', 3.14, true, 1, 1) // Print: Stuff 1 a 3.14 true 1 1
	fmt.Printf("%8f\n", 3.14)                                                 // Eight characters: 3.140000
	fmt.Printf("%.2f\n", 3.141592)                                            // Two decimals: 3.14
	fmt.Printf("%9.f\n", 3.141592)                                            // Eight space in the start: 3
}

func main() {
	MainDataTypes()
	ConvertDataTypes()
	PrintDataTypes()
}
