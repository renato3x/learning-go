package main

import "fmt"

/*
  change swaps the values of two integers using named return values.

  Named returns are return variables that are explicitly declared in the function signature. 
  In this case, `first` and `second` are the names given to the return values.

  Advantages of using named returns:
  - Improve readability by clearly describing the purpose of each returned value.
  - Automatically declare and initialize the return variables at the start of the function.
  - Allow the use of a clean, bare `return` statement without specifying the values again.
  - Enable deferred functions to access and modify return values before the function exits.
*/
func change(p1, p2 int) (first int, second int) {
  first = p2
  second = p1
  return // clean, implicit return (return values of `first` and `second`)
}

var sum = func(value int) func(x int) int {
  return func(x int) int {
    return value + x
  }
}

func main() {
  v1, v2 := change(1, 2)
  fmt.Println(v1, v2)

  adder := sum(5)

  fmt.Println(adder(5))
  fmt.Println(adder(23))
  fmt.Println(adder(17))
}
