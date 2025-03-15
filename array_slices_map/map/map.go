package main

import "fmt"

func map01() {
	// This map is not initialized (current value = nil), which means if you try to set a new key, you'll have an error
	// var approved map[int]string

	approved := make(map[int]string)

	approved[12345678978] = "Maria"
	approved[12345678979] = "Pedro"
	approved[12345678970] = "Ana"
	fmt.Println(approved)

	for cpf, name := range approved {
		fmt.Printf("The student %s, with CPF %d, was approved!\n", name, cpf)
	}

	fmt.Println(approved[12345678970])
	// deletes the element with the specified key (m[key]) from the map
	delete(approved, 12345678970)
	fmt.Println(approved)
}

func map02() {
	employeesAndSalaries := map[string]float64{
		"Jonathan Joestar": 11325.45,
		"Joseph Jostear":   15456.78,
		"Jotaro Kujo":      1200.0,
	}

	employeesAndSalaries["Josuke Higashikata"] = 1350.0
	delete(employeesAndSalaries, "Kars")

	for employee, salary := range employeesAndSalaries {
		fmt.Println(employee, salary)
	}
}

func nestedMap() {
	employeesByLetter := map[string]map[string]float64{
		"J": {
			"Jonathan Joestar": 11325.45,
			"Joseph Jostear":   15456.78,
			"Jotaro Kujo":      1200.00,
		},
		"G": {
			"Giorno Giovanna": 23765.65,
		},
	}

	for letter, employees := range employeesByLetter {
		fmt.Printf("Employees with letter \"%s\":\n\n", letter)

		for employee, salary := range employees {
			fmt.Println(employee, salary)
		}

		fmt.Println()
	}
}

func main() {
	// map01()
	// map02()
	nestedMap()
}
