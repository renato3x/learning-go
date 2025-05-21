package main

import "fmt"

type Car struct {
  name string
  currentSpeed int
}

/* 
  This is a struct that inherits from the Car struct.
  This is a composition, not inheritance.
  The Ferrari struct has a field of type Car, which means it can access the fields and methods of the Car struct.
*/
type Ferrari struct {
  Car
  turbo bool
}

func main() {
  f := Ferrari{}
  f.name = "F40"
  f.currentSpeed = 0
  f.turbo = true

  fmt.Printf("Does the %s have turbo? %t\n", f.name, f.turbo)
}
