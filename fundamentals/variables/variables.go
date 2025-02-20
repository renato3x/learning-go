package main

import (
	"fmt"
	m "math" // "m" is an alias for the "math" package
)

func main() {
  /* 
    -> Defining a constant with an explicit type
    -> constants MUST be initialized
  */
  const PI float64 = 3.1415

  /* 
    -> type type (float64) was inferred (it works to constants too)
    -> if you don't specify the type of the variable, you must to initialize it
    -> if you specify the type of the variable. you don't need to set a value for it on the same line
    -> variables MUST be used
  */
  var radius = 3.2

  /* 
    -> This is a reduced way for creating a var (:=). With this way you don't need to use the "var" keyword to create it,
    but you MUST initialize it
  */
  area := PI * m.Pow(radius, 2)
  fmt.Println("Area of a radius", radius, "is", m.Round(area))

  /* 
    -> This is a way of creating several variables or constants on several lines using the keywords "var"or "const" 
      only once. This way allows you to define different data types to each var/const you create
  */
  const (
    a = 1
    b = 2.2
  )

  var (
    c = "Hello"
    d = false
  )

  fmt.Println(a, b, c, d)

  /* 
    -> This is a way of creating several variables or constants, but in a single line of code. 
    -> When you use this way, the first value goes to the first var/const; the second value goes to the second var/const and so on.
    -> When you use the keywords "var" or "const" you MUST to set a data type for all vars/consts. In other words, all the vars/consts
      will have the same data type
  */
  var e, f bool = true, false
  const g, h string = "Hello", "World!"
  fmt.Println(e, f, g, h)

  /* 
    -> This is a way of creating several variables (not constants) in a single line.
    -> In this case, we're using the reduced way to create vars.
    -> This allows you set different data types for each variable
  */
  i, j, k, l :=  "This is a string", false, 4, 3.4
  fmt.Println(i, j, k, l)
}
