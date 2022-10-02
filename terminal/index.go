package terminal

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
)

func GetUsername() { // Working with the terminal
	fmt.Println("What is your name?")
	reader := bufio.NewReader(os.Stdin)  // Get something from user keyboard
	name, err := reader.ReadString('\n') // Read what the user typed
	if err == nil {                      // Check if has error
		fmt.Println("Hello", name) // Print Hello, name
	} else {
		log.Fatal(err) // Log
	}
}

func Use() {
	// command line: go run main.go 24 43 12 9 10 -- To run with params

	fmt.Println(os.Args) // If you pass values will result: [filename param1 param2] if not pass any param: [filename]
	args := os.Args[1:]  // Slice the args
	iArgs := []int{}

	for _, i := range args {
		val, err := strconv.Atoi(i)
		if err != nil {
			panic(err)
		}
		iArgs = append(iArgs, val)
	}

	max := 0
	for _, val := range iArgs {
		if val > max {
			max = val
		}
	}

	fmt.Println("Max value: ", max)
}
