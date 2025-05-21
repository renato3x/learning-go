package main

import "strings"

type Person struct {
	name     string
	lastName string
}

func (p Person) GetFullName() string {
	return p.name + " " + p.lastName
}

func (p *Person) SetFullName(fullName string) {
	parts := strings.Split(fullName, " ")
	p.name = parts[0]
	p.lastName = parts[1]
}

func main() {
	person := Person{
		name:     "John",
		lastName: "Doe",
	}
	fullName := person.GetFullName()
	println(fullName) // Output: John Doe

	person.SetFullName("Jane Smith")
	fullName = person.GetFullName()
	println(fullName) // Output: Jane Smith
}
