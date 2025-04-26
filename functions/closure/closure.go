package main

import "fmt"

func closure() func() {
  x := 10
  function := func() {
    fmt.Println(x)
  }

  return function
}

func main() {
  x := 20
  fmt.Println(x) // 20

  printX := closure()
  printX() // 10
}
