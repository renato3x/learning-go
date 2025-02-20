package main

import (
	"fmt"
	"math"
	"reflect"
)

func main() {
  // integers
  fmt.Println(1, 2, 1000)

  // unsigned (only positives) uint8 uint16 uint32 uint64
  fmt.Println("Literal integer is", reflect.TypeOf(32000))
  var b byte = 255
  fmt.Println("Byte is", reflect.TypeOf(b))

  i1 := math.MaxInt64

  fmt.Println("The int's maximum value is", i1)
  fmt.Println("Type of i1 is", reflect.TypeOf(i1))

  // string
  text := "Hello, World"
  // len() returns the quantity of characters in a string
  fmt.Println("My string with value \"" + text + "\" have", len(text), "characters")

  // multiple line strings
  text2 := `
    Hello
  
  
    World
  `

  fmt.Println(text2, len(text2))

  char := 'a'

  fmt.Println(reflect.TypeOf(char)) // int32
}
