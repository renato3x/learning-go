package main

import "fmt"

func main() {
  var a int // 0
  var b float64 // 0
  var c bool // false
  var d string // empty string
  var e *int // nil (null) pointer

  fmt.Printf("%v %v %v %q %v\n", a, b, c, d, e)
}
