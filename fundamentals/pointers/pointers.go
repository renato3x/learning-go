package main

import "fmt"

func main() {
  i := 1

  // pointer can receive pointers of type int
  // nil === null
  var pointer *int = nil

  // pointer will receive the memory address from i
  pointer = &i

  // shows the memory address
  fmt.Println(pointer)
  fmt.Println(&i)

  // shows the value stored in the memory address (dereferencing)
  fmt.Println(*pointer)

  *pointer++
  fmt.Println(*pointer)
  fmt.Println(i)
  
  i++
  fmt.Println(*pointer)
  fmt.Println(i)

  *pointer = 5

  fmt.Println(pointer, &i, *pointer, i)
}
