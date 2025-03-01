package main

import (
  "fmt"
  "time"
)

func gradeForConcept(grade float64) string {
  switch {
  case grade >= 9 && grade <= 10:
    return "A"
  case grade >= 8 && grade < 9:
    return "B"
  case grade >= 5 && grade < 8:
    return "C"
  case grade >= 3 && grade < 5:
    return "D"
  default:
    return "E"
  }
}

func main() {
  t := time.Now()

  fmt.Println(gradeForConcept(9.8))
  fmt.Println(gradeForConcept(6.9))
  fmt.Println(gradeForConcept(2.1))

  switch {
  case t.Hour() < 12:
    fmt.Println("Good morning!")
  case t.Hour() < 18:
    fmt.Println("Good afternoon!")
  default:
    fmt.Println("Good night!")
  }
}
