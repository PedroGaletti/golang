package main

import (
	"fmt"
	"math"
	"math/rand"
	"time"
)

func Maths() {
	fmt.Println("\nMaths func: ")
	fmt.Println("5 + 4 = ", 5+4) // Print: 9
	fmt.Println("5 - 4 = ", 5-4) // Print: 1
	fmt.Println("5 * 4 = ", 5*4) // Print: 20
	fmt.Println("5 / 4 = ", 5/4) // Print: 1
	fmt.Println("5 % 4 = ", 5%4) // Print: 1
	mInt := 1
	mInt++
}

func RandomNumbers() {
	fmt.Println("\nRandomNumbers func: ")
	seedSecs := time.Now().Unix()
	fmt.Println("seedSecs", seedSecs)   // Print
	rand.Seed(seedSecs)                 // Seed the seconds
	randNum := rand.Intn(50)            // Generate random number between 0 and 49
	randNum1 := rand.Intn(50) + 1       // Generate random number between 0 and 50
	fmt.Println("Random: ", randNum)    // Print
	fmt.Println("Random 1: ", randNum1) // Print
}

func DefaultMathsFunctions() {
	fmt.Println("\nDefaultMathsFunctions func: ")
	// The Abs function is used to find the absolute value of any given number
	fmt.Println("Abs(-10) = ", math.Abs(-10)) // Abs(-10) = 10
	/*
		With Pow function you can find the base-a exponential of a**b
		Examples: 5**2 = 25
	*/
	fmt.Println("Pow(4, 2) = ", math.Pow(4, 2)) // Pow(4, 2) = 16
	// With Sqrt function you can find the square root of a number
	fmt.Println("Sqrt(16) = ", math.Sqrt(16)) // Sqrt(16) = 4
	/*
		With Cbrt function you can find the cube root of the given input
		Examples: ∛8 = ∛(2 × 2 × 2) = 2
	*/
	fmt.Println("Cbrt(8) = ", math.Cbrt(8)) // Cbrt(8) = 2
	// The Ceil function is used to find the least integer value greater than or equal to the given input
	fmt.Println("Ceil(4.4) = ", math.Ceil(4.4)) // Ceil(4.4) = 5
	// The Floor function is used to find the greatest integer value less than or equal to input
	fmt.Println("Floor(4.4) = ", math.Floor(4.4)) // Floor(4.4) = 4
	// With Round function you can round off the given number to the nearest integer
	fmt.Println("Round(4.4) = ", math.Round(4.4)) // Round(4.4) = 4
	// The Log2 function returns a single value of type int, representing the input argument’s binary logarithm.
	fmt.Println("Log2(8) = ", math.Log2(8)) // Log2(8) = 3
	// The Log10 function is used to find the decimal logarithm of the given input
	fmt.Println("Log10(100) = ", math.Log10(100)) // Log10(100) = 2
	/*
		The Log function is used to find the natural logarithm of the given input
		THe Exp function in used to find the value of e raised to the power equaling the input x
	*/
	fmt.Println("Log(math.Exp(2)) = ", math.Log(math.Exp(2))) // Log(math.Exp(2)) = 2
	// The Max function is used to find the larger of x or y the given inputs
	fmt.Println("Max(5, 4) = ", math.Max(5, 4)) // Max(5, 4) = 5
	// The Min function is used to find the smallest of x or y the given inputs
	fmt.Println("Min(5, 4) = ", math.Min(5, 4)) // Min(5, 4) = 4
}

func Converts() {
	// Convert 90 degrees to radius
	r90 := 90 * math.Pi / 180
	// Convert 90 radius to degrees
	d90 := r90 * (180 / math.Pi)
	fmt.Printf("%.2f radians = %.2f degrees\n", r90, d90)
	fmt.Println("Sin(90) = ", math.Sin(r90)) // Sin(r90) = 1

	/*
		Other maths functions:
		Cos, Tan, Acos, Asin, Atan, Asinh, Acosh,
		Atanh, Atan2, Cosh, Sinh, Sincos Htpot
	*/
}

func main() {
	Maths()
	RandomNumbers()
	DefaultMathsFunctions()
	Converts()
}
