package main

import (
	"fmt"
)

func sum(x int, b int) int {
  return x + b
}

func print(value int) {
  fmt.Println(value)
}

func main() {
  result := sum(3, 4)
  print(result)
}
