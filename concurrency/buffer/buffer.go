package main

import (
	"fmt"
	"time"
)

func routine(ch chan int) {
  // I'm sending to the channel before it is read
  ch <- 1
  ch <- 2
  ch <- 3
  ch <- 4
  fmt.Println("Ran")
  ch <- 5
  ch <- 6
}

func main() {
  ch := make(chan int, 3) // channel with buffer of size 3
  go routine(ch)

  time.Sleep(time.Second)
  fmt.Println(<-ch)
  fmt.Println(<-ch)
  fmt.Println(<-ch)
  fmt.Println(<-ch)
  fmt.Println(<-ch)
}
