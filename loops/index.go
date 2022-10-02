package loops

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

func Loops() {
	// for initialization; condition;
	// postStatement {BODY}
	for x := 1; x <= 5; x++ {
		fmt.Println(x)
	}

	for x := 5; x >= 1; x-- {
		fmt.Println(x)
	}

	fx := 0
	for fx < 5 {
		fmt.Println(fx)
		fx++
	}
}

func While() {
	seedSecs := time.Now().Unix()
	rand.Seed(seedSecs)
	randNum := rand.Intn(51) // Generate a random number
	for true {
		fmt.Print("Guess a number between 0 and 50: ")
		fmt.Println("Random number is: ", randNum)
		reader := bufio.NewReader(os.Stdin)   // Create a new Reader to read the terminal
		guess, err := reader.ReadString('\n') // Read what the user typed in the terminal
		if err != nil {
			log.Fatal(err)
		}

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
	nums := []int{1, 2, 3}
	for _, num := range nums {
		fmt.Println(num)
	}
}
