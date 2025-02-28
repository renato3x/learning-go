package main

import (
  "fmt"
)

func gradeToConcept(n float64) string {
  var grade = int(n)

  switch grade {
  case 10:
    fallthrough
  case 9:
    return "A"
  case 8, 7:
    return "B"
  case 6, 5:
    return "C"
  case 4, 3:
    return "D"
  case 2, 1, 0:
    return "E"
  default:
    return "F"
  }
}

func getFruitPrice(fruit string) {
  switch fruit {
  case "Orange", "Apple":
    fmt.Println("The price from", fruit, "is $", 0.59)
  case "Peach":
    fmt.Println("The price from peach is $", 1.25)
  case "Mango":
    fmt.Println("The price from mango is $", 3.25)
  default:
    fmt.Println("Fruit", fruit, "not found")
  }
}

func main() {
  fmt.Println(gradeToConcept(9.8))
  fmt.Println(gradeToConcept(6.9))
  fmt.Println(gradeToConcept(11.1))
  fmt.Println(gradeToConcept(-1))
  getFruitPrice("Orange")     // The price from Orange is $ 0.59
  getFruitPrice("Apple")      // The price from Apple is $ 0.59
  getFruitPrice("Peach")      // The price from peach is $ 1.25
  getFruitPrice("Mango")      // The price from peach is $ 3.25
  getFruitPrice("Watermelon") // Fruit Watermelon not found
}
