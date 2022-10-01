package maths

import (
	"fmt"
	"math/rand"
	"time"
)

func Maths() {
	fmt.Println("5 + 4 = ", 5+4)
	fmt.Println("5 - 4 = ", 5-4)
	fmt.Println("5 * 4 = ", 5*4)
	fmt.Println("5 / 4 = ", 5/4)
	fmt.Println("5 % 4 = ", 5%4)
	mInt := 1
	mInt++
}

func RandomNumbers() {
	seedSecs := time.Now().Unix()
	fmt.Println("seedSecs", seedSecs)
	rand.Seed(seedSecs)                 // Seed the seconds
	randNum := rand.Intn(50)            // Generate random number between 0 and 49
	randNum1 := rand.Intn(50) + 1       // Generate random number between 0 and 50
	fmt.Println("Random: ", randNum)    // Print
	fmt.Println("Random 1: ", randNum1) // Print
}
