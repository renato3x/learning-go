package main

import "fmt"

func main() {
  s := make([]int, 10)

  fmt.Println(s)
  s[9] = 50
  fmt.Println(s)

  // this create a slice with 10 elements, but the array it references has 20 elements
  s = make([]int, 10, 20)
  fmt.Println(s, len(s), cap(s))

  s = append(s, 1, 2, 3, 4, 5, 6, 7, 8, 9, 0)
  fmt.Println(s, len(s), cap(s))

  s = append(s, 1)
  fmt.Println(s, len(s), cap(s))
}
