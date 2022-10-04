package main

import "fmt"

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
