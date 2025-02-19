# Learning Go
Everything I've studied about Golang

## 1 - Fundamentals

### 1.1 - First

In Go, a package is a way to group related code together. Every Go program is made up of packages. The `main` package is special because it defines the entry point of an executable program.

The "fmt" package provides functions for formatted input and output in Go.

Main Functions:
- `fmt.Println()`: Prints with a newline.
- `fmt.Print()`: Prints without a newline.
- `fmt.Printf()`: Formats and prints using placeholders (e.g., `%s`, `%d`, `%.2f`).
- `fmt.Scan()`, `fmt.Scanln()`, `fmt.Scanf()`: Read user input.

It's commonly used for logging, debugging, and user interaction.

main function is the entrypoint of the entire application

```go
package main

import (
  "fmt"
)

func main() {
  fmt.Println("Hello, world!")
  fmt.Print("Access my github (%s)!", "https://github.com/renato3x")
}
```

### 1.2 - Commands

Go has some basic commands to use in terminal which helps you with somethings

#### `go version`

Show the current version of Go installed in the machine

#### `go env`

Prints Go environment information.

#### `go vet`

Vet examines Go source code and reports suspicious constructs, such as Printf calls whose arguments do not align with the format string. Vet uses heuristics that do not guarantee all reports are genuine problems, but it can find errors not caught by the compilers.

#### `go build`

This command is used to compile Go source code into an executable binary. It analyzes the source files, resolves dependencies, and generates an output file that can be executed.

#### `go run`

Run your Go code
