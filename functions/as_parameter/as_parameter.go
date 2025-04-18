package main

import "fmt"

func multiplication(a, b float64) float64 {
  return a * b
}

func exec(runner func(float64, float64) float64, p1, p2 float64) float64 {
  return runner(p1, p2)
}

func main() {
  result := exec(multiplication, 5.423, 3.823)
  fmt.Println(result)
}
