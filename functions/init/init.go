package main

import "fmt"

/*
  The init function is a special function in Go that is executed before the main function.
  It is often used for initialization tasks, such as setting up global variables or performing setup tasks.

  init function is not required to be defined in the main package, it can be defined in any package.
*/
func init() {
  fmt.Println("Initializing application...")
}

func main() {
  fmt.Println("Application was initialized successfully!")
}
