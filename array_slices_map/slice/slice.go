package main

import (
	"fmt"
	"reflect"
)

func main() {
  /* 
    Slice in Go is a reference to an array. A slice is a lightweight data structure that gives access to a subsequence of the elements of an array.
    Slices do not have a fixed size. They are dynamically-sized. Slices are like references to arrays. A slice does not store any data, it just describes
    a section of an underlying array.
  */

  array1 := [3]int{1, 2, 3}
  slice1 := []int{1, 2, 3}

  fmt.Println(array1, slice1)
  fmt.Println(reflect.TypeOf(array1), reflect.TypeOf(slice1))

  // slice is not an array but a reference to an array
  array2 := [5]int{1, 2, 3, 4, 5}
  slice2 := array2[1:3] // (2, 3) new slice but points to array2, so changes in slice2 will reflect in array2
  fmt.Println(array2, slice2)

  slice3 := array2[:2] // new slice but points to array2, so changes in slice3 will reflect in array2
  fmt.Println(array2, slice3)

  slice4 := slice2[:1] // implicitly starts from 0, so slice4 is (2)
  fmt.Println(array2, slice4)
}
