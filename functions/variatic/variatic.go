package main

import "fmt"

// this function can receive a variable number of arguments
func average(grades...float64) (result float64) {
  sum := 0.0

  for _, value := range grades {
    sum += value
  }

  result = sum / float64(len(grades))

  return
}

func main() {
  fmt.Println(average(1, 2, 3, 4, 5, 6, 7, 8, 9, 9))
}
