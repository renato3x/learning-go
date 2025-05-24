package main

import "fmt"

type SportsCar interface {
  turnOnTurbo()
}

type Ferrari struct {
  model string
  turbo bool
  currentSpeed int
}

func (f *Ferrari) turnOnTurbo() {
  fmt.Println(f)
  f.turbo = true
}

func main() {
  car := Ferrari{
    model: "Ferrari 488",
    turbo: false,
    currentSpeed: 0,
  }
  car.turnOnTurbo()

  // We must use a pointer to Ferrari here because the method turnOnTurbo()
  // has a pointer receiver (*Ferrari). Only *Ferrari implements the SportsCar interface,
  // not Ferrari (the value type), since in Go, *Ferrari and Ferrari are distinct types.
  var car2 SportsCar = &Ferrari{
    model: "Ferrari 488",
    turbo: false,
    currentSpeed: 0,
  }
  
  fmt.Println(car, car2)
}
