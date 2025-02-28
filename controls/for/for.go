package main

import (
  "fmt"
  "time"
)

func main() {
  i := 1

  for i <= 10 { // similar to while
    fmt.Printf("%d ", i)
    i++
  }
  fmt.Println()

  for i := 0; i <= 20; i += 2 { // traditional for
    fmt.Printf("%d ", i)
    time.Sleep(time.Second * 3)
  }
  fmt.Println()

  for i := 1; i <= 6; i++ {
    if i%2 == 0 {
      fmt.Printf("Even, ")
    } else {
      fmt.Printf("Odd, ")
    }
  }
  fmt.Println()
}
