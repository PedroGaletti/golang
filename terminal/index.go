package terminal

import (
	"bufio"
	"fmt"
	"log"
	"os"
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
