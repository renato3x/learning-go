package main

import (
	"fmt"
	"time"
)

func talk(person string) <-chan string {
  c := make(chan string)

  go func() {
    for i := range 3 {
      time.Sleep(time.Second)
      c <- fmt.Sprintf("%s talking %d", person, i)
    }

    close(c)
  }()

  return c
}

func merge(input1, input2 <-chan string) <-chan string {
  c := make(chan string)

  go func() {
    for {
      select {
      case msg := <-input1:
        c <- msg
      case msg := <-input2:
        c <- msg
      }
    }
  }()

  return c
}

func main() {
  c := merge(talk("Alice"), talk("Bob"))
  fmt.Println(<-c, <-c)
  fmt.Println(<-c, <-c)
  fmt.Println(<-c, <-c)
}
