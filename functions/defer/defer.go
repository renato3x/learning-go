package main

import "fmt"

func main() {
  fmt.Println(getApprovedValue(3500))
}

func getApprovedValue(number int) int {
  /* 
    defer is a keyword in Go that allows you to schedule a function call to be run after the function completes.
    It is often used for cleanup tasks, such as closing files or releasing resources.
    The deferred function call is executed in the reverse order of their declaration.
    In this case, the deferred function will be executed after the return statement.
  */
  defer fmt.Println("end!")
  
  if number >= 5000 {
    fmt.Println("High value...")
    return number
  }

  fmt.Println("Low value...")
  return number
}
