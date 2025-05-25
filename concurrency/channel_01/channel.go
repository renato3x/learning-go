package main

import "fmt"

func main() {
	ch := make(chan any, 1)

	ch <- 1 // Send 1 to the channel
	<-ch    // Receive from the channel and discard it

	ch <- 2           // Send 2 to the channel
	fmt.Println(<-ch) // Receive from the channel and print it
}
