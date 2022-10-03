package main

import (
	"fmt"
	"math/rand"
	"time"
)

func Maths() {
	fmt.Println("\nMaths func:")
	fmt.Println("5 + 4 = ", 5+4) // Print: 9
	fmt.Println("5 - 4 = ", 5-4) // Print: 1
	fmt.Println("5 * 4 = ", 5*4) // Print: 20
	fmt.Println("5 / 4 = ", 5/4) // Print: 1
	fmt.Println("5 % 4 = ", 5%4) // Print: 1
	mInt := 1
	mInt++
}

func RandomNumbers() {
	fmt.Println("\nRandomNumbers func:")
	seedSecs := time.Now().Unix()
	fmt.Println("seedSecs", seedSecs)   // Print
	rand.Seed(seedSecs)                 // Seed the seconds
	randNum := rand.Intn(50)            // Generate random number between 0 and 49
	randNum1 := rand.Intn(50) + 1       // Generate random number between 0 and 50
	fmt.Println("Random: ", randNum)    // Print
	fmt.Println("Random 1: ", randNum1) // Print
}

func main() {
	Maths()
	RandomNumbers()
}
