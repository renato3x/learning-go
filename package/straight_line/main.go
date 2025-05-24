package main

import "fmt"

func main() {
	p1, p2 := Point{2, 2}, Point{2, 4}
	fmt.Println(legs(p1, p2))
	fmt.Println(Distance(p1, p2))
}
