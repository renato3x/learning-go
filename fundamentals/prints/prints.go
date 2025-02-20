package main

import "fmt"

func main() {
  fmt.Print("Same")
  fmt.Print(" line.")
  fmt.Println(" New")
  fmt.Println("Line.")
  x := 3.141516

  // fmt.Println("The value of x is " + x) (error)

  xs := fmt.Sprint(x)
  fmt.Println("The value of x is " + xs)

  fmt.Println("The value of x is", x)
  fmt.Printf("The value of x is %f\n", x)

  a := 1
  b := 1.9999
  c := false
  d := "Hi"

  fmt.Printf("%d %f %.1f %t %s\n", a, b, b, c, d)
  // %v is a generic hook that works to any values
  fmt.Printf("%v %v %v %v %v\n", a, b, b, c, d)
}
