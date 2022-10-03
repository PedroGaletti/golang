package main

import (
	"bufio"
	"fmt"
	"log"
	"math/rand"
	"os"
	"strconv"
	"strings"
	"time"
)

// for initialization; condition;
// postStatement {BODY}

func Loops() {
	fmt.Println("\nLoops func:")
	for x := 1; x <= 5; x++ {
		fmt.Print(x) // Print: 1 2 3 4 5
	}

	fmt.Printf("\n") // Just to separate the logs

	for x := 5; x >= 1; x-- {
		fmt.Print(x) // Print: 5 4 3 2 1
	}

	fmt.Printf("\n") // Just to separate the logs

	fx := 0
	for fx < 5 {
		fmt.Print(fx) // Print: 0 1 2 3 4
		fx++
	}
}

func While() {
	fmt.Println("\nWhile func:")
	seedSecs := time.Now().Unix()
	rand.Seed(seedSecs)
	randNum := rand.Intn(51) // Generate a random number
	for true {
		fmt.Print("Guess a number between 0 and 50: \n")
		reader := bufio.NewReader(os.Stdin)   // Create a new Reader to read the terminal
		guess, err := reader.ReadString('\n') // Read what the user typed in the terminal
		if err != nil {
			log.Fatal(err)
		}

		fmt.Println("Random number is: ", randNum)

		guess = strings.TrimSpace(guess)   // Remove all the spaces in the start and the end
		iGuess, err := strconv.Atoi(guess) // Converts the string into int
		if err != nil {
			log.Fatal(err)
		}
		if iGuess > randNum {
			fmt.Println("Pick a lower value")
		} else if iGuess < randNum {
			fmt.Println("Pick a higher value")
		} else {
			fmt.Println("You guessed it")
			break // **Break** jump to the beggin and **continue** jump to the first line of loop
		}
	}
}

func Ranges() {
	fmt.Println("\nRanges func:")
	nums := []int{1, 2, 3}
	for _, num := range nums {
		fmt.Print(num) // Print: 1 2 3
	}

	fmt.Printf("\n") // Just to separate the logs
}

func main() {
	Loops()
	While()
	Ranges()
}
