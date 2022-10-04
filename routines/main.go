package main

import (
	"fmt"
	"time"
)

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
