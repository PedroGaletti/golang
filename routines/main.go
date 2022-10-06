package main

import (
	"fmt"
	"time"
)

/*
GO routines is a function or method that executes concurrently alongside any other goroutines using a special goroutine thread.

Goroutine thread are more lightweight than standart threads, with most Golang programs using thousands of goroutines at once.

TO create a goroutine, add the go keyword before the function declaration.

You can stop a goroutine by sending it a signal channel. Goroutines can only respond to signals if told to check, so you'll need to include checks in logical places such as the top of your loop.

all the goroutines terminate when the main function terminates.
*/

func FirstRoutine() {
	for i := 1; i <= 15; i++ {
		fmt.Println("Func 1: ", i)
	}
}

func SecondRoutine() {
	for i := 1; i <= 10; i++ {
		fmt.Println("Func 2: ", i)
	}
}

func main() {
	go FirstRoutine()
	go SecondRoutine()
	time.Sleep(2 * time.Second)
}
