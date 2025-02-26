package main

import (
  "fmt"
  "math/rand"
  "time"
)

func printResult(grade float64) {
  if grade >= 7 {
    fmt.Println("Approved with note", grade)
  } else {
    fmt.Println("Not approved with note", grade)
  }
}

func gradeForConcept(grade float64) string {
  if grade >= 9 && grade <= 10 {
    return "A"
  } else if grade >= 8 && grade < 9 {
    return "B"
  } else if grade >= 5 && grade < 8 {
    return "C"
  } else if grade >= 3 && grade < 5 {
    return "D"
  } else {
    return "E"
  }
}

func randomNumber() int {
  s := rand.NewSource(time.Now().UnixNano())
  r := rand.New(s)
  return r.Intn(10)
}

func main() {
  // if/else
  printResult(7.3)
  printResult(5.1)

  // if/else-if/else
  fmt.Println(gradeForConcept(9.8))
  fmt.Println(gradeForConcept(6.9))
  fmt.Println(gradeForConcept(2.1))

  // if init
  /*
     “i” is a variable that can be accessed within the scope of if/else.
     Outside of it, it can no longer be accessed, just like in a for loop
  */
  if i := randomNumber(); i > 5 {
    fmt.Println("Win :)")
  } else if i == 5 {
    fmt.Println("Almost win :)")
  } else {
    fmt.Println("Lose :(")
  }
}
