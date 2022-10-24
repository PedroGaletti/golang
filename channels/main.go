package main

import "fmt"

/*
Channels are the medium through which goroutines can communicate with each other.
In simple terms, a channel is a pipe that allows a goroutine to either put or read the data.
*/

func FirstChannels(channel chan int) {
	channel <- 1
	channel <- 2
	channel <- 3
}

func SecondChannels(channel chan int) {
	channel <- 4
	channel <- 5
	channel <- 6
}

func main() {
	var c chan int
	fmt.Println(c)                         // nil
	fmt.Printf("Type of channel: %T\n", c) // Type of channel: chan int

	fChannel := make(chan int)
	sChannel := make(chan int)

	go FirstChannels(fChannel)
	go SecondChannels(sChannel)

	fmt.Println("<-fChannel: ", <-fChannel)
	fmt.Println("<-fChannel: ", <-fChannel)
	fmt.Println("<-fChannel: ", <-fChannel)
	fmt.Println("<-sChannel: ", <-sChannel)
	fmt.Println("<-sChannel: ", <-sChannel)
	fmt.Println("<-sChannel: ", <-sChannel)
}
