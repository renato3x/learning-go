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

### 1.3 - Variables

In Golang, the variables have static types, which means the variables will have one data type until the and of the code, different of JavaScript or Python, for example. When you create a variable in Go, you can the type of the variable explicitly, or the language will infer the type automatically

Go has two keywords for creating variables: `var` and `const`. When you use `var`, the variable can be given an initial value and you can change this value at any time. But, when you use `const`, the variable must be initialized with some value and this value cannot be changed anymore.

See some examples below:

```go
package main

import "fmt"

func main() {
  // creating normal variable with explicit type "int"
  var age int = 21

  // creating a constant value with explicit type "string"
  const name string = "Renato"

  fmt.Printf("My name is %s and I have %d years old", name, age)
}
```

```go
package main

import "fmt"

func main() {
  // creating normal variable with inferred type "int"
  var age = 21

  // creating a constant value with inferred type "string"
  const name = "Renato"

  fmt.Printf("My name is %s and I have %d years old", name, age)
}
```

When using regular variables, you can use another syntax that simplifies the code: an operator called **short variable declaration**. With this operator, you can declare a variable without the `var` keyword, as shown in the example below:

```go
package main

import "fmt"

func main() {
  // creating normal variable using "short variable declaration" operator
  age := 21

  fmt.Printf("I have %d years old", age)
}
```

This operator allows you to declare a variable without using `var` keyword, but you must initialize it with a value. It also automatically infers the variable’s type.

Just like in other languages you can also declare multiple variable in a single line of code. See the examples below:

```go
package main

import "fmt"

func main() {
  // declaring multiple constants of type float64 in a single line
  const a, b float64 = 1.1, 2.2

  // declaring multiple variables of type string in a single line
  var c, d string = "Hello", "world!"

  fmt.Println(a, b)
  fmt.Println(c, d)
}
```

When you declare multiple variables or constants in a single line with an explicit type, all of them will have the same data type. However, if no type is specified, Go will infer the type of each value individually.

```go
package main

import "fmt"

func main() {
  // const a with a float64 value and const b with a boolean value
  const a, b = 1.2, false

  // var c with a int value and var d with a string value
  var c, d = 1, "world!"

  fmt.Println(a, b)
  fmt.Println(c, d)
}
```

When you declare multiple variables in a single line you can also use the short variable declaration operator. Using this operator in this context lets you assign different data types to each variable too

```go
package main

import "fmt"

func main() {
  a, b :=  "I love you", false
  fmt.Println(a, b)
}
```

You can also declare multiple variables or constants using multiple lines, and this method allows you to declare variables and constants of any type

```go
package main

import "fmt"

func main() {
  // declaring multiple constants across multiple lines, assigning different data types to each
  const (
    a = 1.1
    b = true
  )

  // declaring multiple variables across multiple lines, assigning different data types to each
  var (
    c = 21
    d = "Hello"
  )

  fmt.Println(a, b)
  fmt.Println(c, d)
}
```

**DO NOT FORGET**: Regular variables **must be used in your code!**. If you do not use the regular variables you declare, your code will not compile

### 1.4 - Prints

Nothing relevant to say here...

### 1.5 - Types

Nothing relevant to say here...

### 1.6 - "Zeros"

In Go, the basic types has the following initial types if you don't initialize with a specific value:

```go
package main

func main() {
  var a int // 0
  var b float64 // 0
  var c bool // false
  var d string // empty string
  var e *int // nil (null) pointer
}
```
