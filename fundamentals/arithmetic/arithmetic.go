package main

import (
	"fmt"
	"math"
)

func main() {
  a := 3
  b := 2

  fmt.Println("Sum ->", a + b)
  fmt.Println("Subtraction ->", a - b)
  fmt.Println("Division ->", a / b)
  fmt.Println("Multiplication ->", a / b)
  fmt.Println("Module ->", a % b)

  // bitwise
  fmt.Println("And ->", a & b)
  fmt.Println("Or ->", a | b)
  fmt.Println("Xor ->", a ^ b)

  // using math
  c := 3.0
  d := 2.0

  fmt.Println("Greater -> ", math.Max(c, d))
  fmt.Println("Less -> ", math.Min(c, d))
  fmt.Println("Power -> ", math.Pow(c, d))
}
