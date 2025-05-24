package main

import "fmt"

/* 
  Custom types can "extend" the functionality of built-in types.
  In this example, we create a custom type called "Grade" that is based on the built-in float64 type.
*/

type Grade float64

func (g Grade) between(start, end float64) bool {
  return float64(g) >= start && float64(g) <= end
}

func gradeToConcept(grade Grade) string {
  switch {
  case grade.between(9, 10):
    return "A"
  case grade.between(7, 8.9):
    return "B"
  case grade.between(5, 7.9):
    return "C"
  case grade.between(3, 4.9):
    return "D"
  default:
    return "E"
  }
}

func main() {
  fmt.Println(gradeToConcept(9.5)) // Output: A
  fmt.Println(gradeToConcept(8.0)) // Output: B
  fmt.Println(gradeToConcept(6.5)) // Output: C
  fmt.Println(gradeToConcept(4.5)) // Output: D
  fmt.Println(gradeToConcept(2.0)) // Output: E
}
