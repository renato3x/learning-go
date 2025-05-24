package main

import (
	"fmt"
)

type SportsCar interface {
  turnOnTurbo()
}

type Deluxe interface {
  parallelPark()
}

type DeluxeSportsCar interface {
  SportsCar
  Deluxe
  // you can add more interfaces or specific methods here
}

type BMW7 struct {}

func (b BMW7) parallelPark() {
  fmt.Println("TURBO ON");
}

func (b BMW7) turnOnTurbo() {
  fmt.Println("Parallel parking");
}

func main() {
  var car DeluxeSportsCar = BMW7{}
  car.parallelPark()
  car.turnOnTurbo()
  fmt.Println(car)
}
