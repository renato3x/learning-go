package main

import "fmt"

type Course struct {
  name string
}

func main() {
  /* 
    An empty interface can hold any value.
    The empty interface is the most general type in Go.
    It is defined as interface{}.
    The empty interface is used when you don't know the type of the value you are going to store.
  */
  var thing interface{}
  fmt.Println(thing)

  thing = 3
  fmt.Println(thing)
  thing = "hello"
  fmt.Println(thing)
  thing = Course{"Go"}
  fmt.Println(thing)
}
