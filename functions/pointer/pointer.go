package main

import "fmt"

func main() {
	number := 1
	numberPointer := &number

	increment(number)
	fmt.Println(number)

	incrementPointer(numberPointer)
	fmt.Println(number)
}

func increment(number int) {
	number++
}

func incrementPointer(number *int) {
	*number++
}
