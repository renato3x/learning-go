package main

import "fmt"

func typeof(i interface{}) string {
  switch i.(type) {
  case int:
    return "int"
  case float32, float64:
    return "float"
  case string:
    return "string"
  case func():
    return "func"
  default:
    return "unknown"
  }
}

func main() {
	fmt.Println(typeof(5))
  fmt.Println(typeof(5.3))
  fmt.Println(typeof(float32(4.5)))
  fmt.Println(typeof("Hello World"))
  fmt.Println(typeof(func() {}))
  fmt.Println(typeof(nil))
}
