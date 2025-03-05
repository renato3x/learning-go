package main

import "fmt"

func main() {
	numbers := [...]int{1, 2, 3, 4, 5}

  for index, number := range numbers {
    fmt.Printf("%d) %d\n", index, number)
  }

  // ignoring the current index
  for _, number := range numbers {
    fmt.Printf("%d\n", number)
  }
}
