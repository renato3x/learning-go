package main

import (
	"fmt"
)

func main() {
  result := factorial(5)
  fmt.Printf("Factorial of %d is %d\n", 5, result)
}

func factorial(number uint) uint {
  if (number == 0 || number == 1) {
    return 1
  }

  return number * factorial(number - 1)
}
