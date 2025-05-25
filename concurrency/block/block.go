package main

import (
	"fmt"
	"time"
)

func routine(ch chan int) {
  time.Sleep(time.Second);
  ch <- 1
  fmt.Println("Only after it was read")
}

func main() {
  ch := make(chan int) // channel without buffer
  go routine(ch)

  fmt.Println(<-ch)
  fmt.Println("Was read")
  fmt.Println(<-ch) // deadlock, because the channel is unbuffered and no other goroutine is sending to it
  fmt.Println("End")
}
