package main

import "fmt"

/*
Interfaces are implemented implicitly in Go.
An interface is a type that specifies a contract for other types to implement.
An interface defines a set of methods that a type must implement to satisfy the interface.
An interface can be used to define a common behavior for different types.
*/
type Printable interface {
	toString() string
}

type Person struct {
	name    string
	surname string
}

type Product struct {
	name  string
	price float64
}

func (p Person) toString() string {
	return p.name + " " + p.surname
}

func (p Product) toString() string {
	return p.name + " " + fmt.Sprintf("%.2f", p.price)
}

func print(p Printable) {
	fmt.Println(p.toString())
}

func main() {
	print(Person{"John", "Doe"})
	print(Product{"Laptop", 999.99})
}
