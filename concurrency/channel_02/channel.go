package main

import (
	"fmt"
	"time"
)

func twoThreeFourTimes(base int, ch chan int) {
	time.Sleep(time.Second)
	ch <- 2 * base

	time.Sleep(time.Second)
	ch <- 3 * base

	time.Sleep(time.Second)
	ch <- 4 * base
}

func main() {
	ch := make(chan int)
	go twoThreeFourTimes(2, ch)
  fmt.Println("A")

	a, b := <-ch, <-ch // Receive values from the channel
  fmt.Println(a, b)

  fmt.Println(<-ch)
}
