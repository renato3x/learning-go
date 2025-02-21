package main

import (
	"fmt"
	"strconv"
)

func main() {
  x := 2.4
  y := 2

  fmt.Println(x / float64(y))

  grade := 6.9
  finalGrade := int(grade)

  fmt.Println(finalGrade)

  // careful...
  fmt.Println("Test " + string(97)) // Test a

  // int to string
  // Itoa -> Integer to ASCII
  fmt.Println("Test " + strconv.Itoa(97)) // Test 97

  // string to int
  // Atoi -> ASCII to Integer
  num, _ := strconv.Atoi("97")
  fmt.Println(num - 122)

  // string to bool
  boolean, _ := strconv.ParseBool("true")

  if boolean {
    fmt.Println("True")
  } else {
    fmt.Println("False")
  }

  // string to float
  float_value, _ := strconv.ParseFloat("1.234", 64)
  fmt.Println(float_value)
}
