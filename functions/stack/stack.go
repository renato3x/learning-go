package main

import (
	"fmt"
	"runtime/debug"
)

func func3() {
	debug.PrintStack()
}

func func2() {
	func3()
}

func func1() {
	func2()
}

func mapper[T, R any](items []T, mapfn func(value T, index int) R) []R {
	final := make([]R, len(items))

	for index, value := range items {
		final[index] = mapfn(value, index)
	}

	return final
}

func main() {
	func1()

	// the code below has nothing to do with the content from this class
	phrases := []string{
		"Today is a great day.",
		"I love coding in Go!",
		"Coffee makes mornings better.",
		"Always keep learning.",
		"Music brings joy to life.",
		"Pizza is my favorite food.",
		"Stay focused and productive.",
		"Reading expands your mind.",
		"Organization leads to success.",
		"Life is full of surprises.",
	}

	getLengths := func(text string, _ int) int {
		return len(text)
	}

	final := mapper(phrases, getLengths)

	fmt.Println(phrases)
	fmt.Println(final)
}
