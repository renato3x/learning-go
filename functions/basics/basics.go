package main

import "fmt"

func func1() {
	fmt.Println("func1")
}

func func2(p1 string, p2 string) {
	fmt.Printf("func2: %s %s \n", p1, p2)
}

func func3() string {
	return "func3"
}

// both params are strings
func func4(p1, p2 string) string {
	return fmt.Sprintf("func4: %s %s", p1, p2)
}

func func5() (string, string) {
	return "Hello", "World"
}

func main() {
  func1()
  func2("Hello", "World")
  fmt.Println(func3())
  fmt.Println(func4("Hello", "World"));

  v1, v2 := func5()

  fmt.Println(v1, v2)
}
