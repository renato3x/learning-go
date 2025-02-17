/*
	In Go, a package is a way to group related code together.
	Every Go program is made up of packages.
	The `main` package is special because it defines the entry point of an executable program.
*/
package main

/*
	The "fmt" package provides functions for formatted input and output in Go.

	Main Functions:
	- `fmt.Println()`: Prints with a newline.
	- `fmt.Print()`: Prints without a newline.
	- `fmt.Printf()`: Formats and prints using placeholders (e.g., `%s`, `%d`, `%.2f`).
	- `fmt.Scan()`, `fmt.Scanln()`, `fmt.Scanf()`: Read user input.

	It's commonly used for logging, debugging, and user interaction.
*/

import "fmt"

// main function is the entrypoint of the entire application
func main() {
	fmt.Println("Hello, world!")
	fmt.Println("Hello, Golang!")
}
