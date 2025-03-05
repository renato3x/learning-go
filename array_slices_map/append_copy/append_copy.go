package main

import "fmt"

func main() {
  array := [3]int{1, 2, 3}
  var slice []int

  // array = append(array, 4, 5, 6) (not possible)
  slice = append(slice, 4, 5, 6)
  fmt.Println(array, slice)

  slice2 := make([]int, 2)
  copy(slice2, slice)
  // as slice2 is of length 2, only first 2 elements of slice are copied (4, 5).
  // copy do not increase the length of slice2, just copy the exact number of elements that can fit in slice2
  fmt.Println(slice2)
}
