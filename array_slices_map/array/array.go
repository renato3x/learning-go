package main

import "fmt"

func forEach(items []int, consumer func(value int)) {
  for _, item := range items {
    consumer(item)
  }
}

func main() {
	// declaring a array with 3 positions of type float64
	var grades [3]float64
	fmt.Println(grades) // [0 0 0]

	grades[0], grades[1], grades[2] = 7.8, 4.3, 9.1

	fmt.Println(grades) // [7.8 4.3 9.1]

	total := 0.0

	// This is not the best practice to loop an array
	for i := 0; i < len(grades); i++ {
		total += grades[i]
	}

	average := total / float64(len(grades))

	fmt.Printf("Average %.2f\n", average)
}
