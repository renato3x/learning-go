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

### 1.7 - Conversions

There are many ways to convert types in Go. Here are some ways to convert some types

#### Integer to Float64

```go
package main

import (
  "fmt"
  "reflect" 
)

func main() {
  x := 6
  parsed_x := float64(x)
  fmt.Println(reflect.TypeOf(parsed_x)) // float64
}
```

#### Float64 to Integer

When you parse a float into an integer, the value will be truncated instead of being rounded

```go
package main

import (
  "fmt"
  "reflect" 
)

func main() {
  x := 6.5
  parsed_x := int(x)
  fmt.Println(parsed_x) // 6
  fmt.Println(reflect.TypeOf(parsed_x)) // int64
}
```

#### Integer to string

Based on the last two examples, the most logical is to use `string()` to parse an integer to string, but...

```go
package main

import (
  "fmt"
)

func main() {
  x := 97
  parsed_x := string(x)
  fmt.Println(parsed_x) // a
}
```

Why the output in console is "a"? This method will transform the number into the character it corresponds to in the ASCII table. In this case, 97 corresponds to the lowercase "a" in the ASCII table

To really parse an integer to a string, you can use the package `strconv` do do this.

```go
package main

import (
  "fmt"
  "strconv"
  "reflect"
)

func main() {
  x := 5
  parsed_x := strconv.Itoa(x)
  fmt.Println(parsed_x) // 5
  fmt.Println(reflect.TypeOf(parsed_x)) // string
}
```

the function `Itoa` (Integer to ASCII) literally parse the integer value to a string. And you can also make the reverse process, using `Atoi`

```go
package main

import (
  "fmt"
  "strconv"
  "reflect"
)

func main() {
  x := "5"
  parsed_x, _ := strconv.Atoi(x)
  fmt.Println(parsed_x) // 5
  fmt.Println(reflect.TypeOf(parsed_x)) // int
}
```

This function has to returns: the parsed value and an error, respectively. But, the error is being ignored using the underscore symbol (_). Later we'll be errors better. And yes, **functions can return more than one value**.

#### String to boolean

Basically, when you try to parse a string to a boolean, the values "true" and "1" are considered as the boolean true. Values different of these two will be considered the boolean false. To parse a string to boolean we use the package strconv too, using the function `ParseBool`

```go
package main

import (
  "fmt"
  "strconv"
  "reflect"
)

func main() {
  a, _ := strconv.ParseBool("true")
  fmt.Println(a) // true

  b, _ := strconv.ParseBool("1")
  fmt.Println(b) // true

  c, _ := strconv.ParseBool("false")
  fmt.Println(c) // false

  d, _ := strconv.ParseBool("0")
  fmt.Println(d) // false

  e, _ := strconv.ParseBool("-1")
  fmt.Println(e) // false

  f, _ := strconv.ParseBool("2")
  fmt.Println(f) // false

  g, _ := strconv.ParseBool("Hello World")
  fmt.Println(g) // false
}
```

Like `Atoi`, the ParseBool function returns the parsed value and an error if there is.

#### String to float

To parse string to float we still use the package strconv, but we use the function `ParseFloat`, which follows the pattern as `Atoi` and `ParseBool`,  but it also requires the bitSize to be passed in (32 or 64)

```go
package main

import (
  "fmt"
  "strconv"
  "reflect"
)

func main() {
  x := "5.574"
  parsed_x, _ := strconv.ParseFloat(x, 64)
  fmt.Println(parsed_x) // 5.574
  fmt.Println(reflect.TypeOf(parsed_x)) // float64
}
```

### 1.8 - Functions (Basic)

Nothing special...

### 1.9 - Arithmetic

Nothing special...

### 1.10 - Assignment

Nothing special...

### 1.11 - Relational

Nothing special...

### 1.12 - Logical

Nothing special...

### 1.12 - Unary

Nothing special...

### [1.13 - Pointers](fundamentals/pointers/pointers.go)

## 2 - Control Structures

### 2.1 - if/else

Just like in a for loop, you can create a variable that can be accessed only within the scope of if/else or if/else-if/else, as the example below:

```go
package main

import (
  "fmt"
  "rand"
  "time"
)

func generateAge() int {
  s := rand.NewSource(time.Now().UnixNano())
  r := rand.New(s)
  return r.Intn(30)
}

func main() {
  if age := generateAge(); age >= 18 {
    fmt.Println("You're of legal age")
  } else if age >= 16 {
    fmt.Println("You're almost of age")
  } else {
    fmt.Println("You're a long way from the age of majority")
  }
}
```

### 2.2 - For (Basic)

Nothing special to comment...

### 2.3 - Switch/Case

The Switch/Case structure has the same idea as other languages. But, instead of the Switch keeping checking all the
cases even if one has already been answered, when a case is answered, it executes it and stops executing the Switch.

```go
package main

import "fmt"

func getFruitPrice(fruit string) {
	switch fruit {
  case "Orange":
    fmt.Println("The price from orange is $", 0.59)
  case "Apple":
    fmt.Println("The price from apple is $", 0.59)
  case "Peach":
    fmt.Println("The price from peach is $", 1.25)
  case "Mango":
    fmt.Println("The price from mango is $", 3.25)
  default:
    fmt.Println("Fruit", fruit, "not found")
	}
}

func main() {
  getFruitPrice("Orange") // The price from orange is $ 0.59
  getFruitPrice("Apple") // The price from apple is $ 0.59
  getFruitPrice("Peach") // The price from peach is $ 1.25
  getFruitPrice("Mango") // The price from peach is $ 3.25
  getFruitPrice("Watermelon") // Fruit Watermelon not found
}
```

As you can see in the example above, the fruits "Orange" and "Apple" have the same price and show a similar output. So,
we can simplify this code by using ways that simulate the default switch/case functionality from other languages, which
is executing multiple cases even if one has already been answered:

You can use the `fallthrough` keyword

```go
package main

import "fmt"

func getFruitPrice(fruit string) {
	switch fruit {
  case "Orange":
    fallthrough
  case "Apple":
    fmt.Println("The price from", fruit ,"is $", 0.59)
  case "Peach":
    fmt.Println("The price from peach is $", 1.25)
  case "Mango":
    fmt.Println("The price from mango is $", 3.25)
  default:
    fmt.Println("Fruit", fruit, "not found")
	}
}

func main() {
  getFruitPrice("Orange") // The price from Orange is $ 0.59
  getFruitPrice("Apple") // The price from Apple is $ 0.59
  getFruitPrice("Peach") // The price from peach is $ 1.25
  getFruitPrice("Mango") // The price from peach is $ 3.25
  getFruitPrice("Watermelon") // Fruit Watermelon not found
}
```

Or you can separate the values using `,`:

```go
package main

import "fmt"

func getFruitPrice(fruit string) {
	switch fruit {
  case "Orange", "Apple":
    fmt.Println("The price from", fruit ,"is $", 0.59)
  case "Peach":
    fmt.Println("The price from peach is $", 1.25)
  case "Mango":
    fmt.Println("The price from mango is $", 3.25)
  default:
    fmt.Println("Fruit", fruit, "not found")
	}
}

func main() {
  getFruitPrice("Orange") // The price from Orange is $ 0.59
  getFruitPrice("Apple") // The price from Apple is $ 0.59
  getFruitPrice("Peach") // The price from peach is $ 1.25
  getFruitPrice("Mango") // The price from peach is $ 3.25
  getFruitPrice("Watermelon") // Fruit Watermelon not found
}
```

You can also have a similar functionality from an if/else inside the switch. If you don't specify a variable or value to
compare in the switch, Go will automatically understand comparison value as `true`, which lets you create cases 
within the switch that have boolean tests. The first `true` case will be executed

```go
package main

import "fmt"

func checkAge(age int) {
  switch {
  case age < 18:
    fmt.Println("You are a young")
  case age < 50:
    fmt.Println("You are an adult")
  default:
    fmt.Println("You are an old")
  }
}

func main() {
  checkAge(17) // You are a young
  checkAge(18) // You are an adult
  checkAge(12) // You are a young
  checkAge(20) // You are an adult
  checkAge(50) // You are an old
  checkAge(52) // You are an old
}
```


## 3 - Arrays, Slices & Maps

### 3.1 - Arrays

In Go, an array is a fixed-length sequence of elements of a single type. The length of an array is part of its type, which means that arrays cannot be resized. Here is an example of how to declare and use an array in Go:

```go
package main

import "fmt"

func main() {
  // declaring an array with 4 positions of type int
  var numbers [4]int
  fmt.Println(numbers) // [0 0 0 0]

  numbers[0], numbers[1], numbers[2], numbers[3] = 10, 20, 30, 40

  fmt.Println(numbers) // [10 20 30 40]

  sum := 0

  // Looping through the array using a for loop
  for i := 0; i < len(numbers); i++ {
    sum += numbers[i]
  }

  fmt.Printf("Sum: %d\n", sum)
}
```

In the example above:
- We declare an array `numbers` with 4 positions of type `int`.
- We assign values to each position in the array.
- We calculate the sum of the numbers using a `for` loop.

### 3.2 - For Range

The `for range` loop in Go is a powerful construct that allows you to iterate over elements in various data structures, such as arrays, slices, maps, and strings. Here is an example of how to use the `for range` loop with an array:

```go
package main

import "fmt"

func main() {
  letters := [...]string{"A", "B", "C", "D"}

  for index, letter := range letters {
    fmt.Printf("%d) %s\n", index, letter)
  }

  // ignoring the current index
  for _, letter := range letters {
    fmt.Printf("%s\n", letter)
  }
}
```

In the example above:
- We declare an array `letters` with 4 positions of type `string` using the `[...]` syntax, which allows the compiler to determine the length of the array based on the number of elements provided.
- We use the `for range` loop to iterate over the array, printing both the index and the value of each element.
- We use the `for range` loop again to iterate over the array, this time ignoring the index and printing only the values.

The `for range` loop is a convenient way to iterate over elements in a data structure without needing to manage the loop counter manually. The `range` keyword returns both the index and the value of each element in the array, which can be useful for various operations.

### 3.3 - Slices

Slices in Go are a reference to an underlying array. They provide a more flexible and dynamic way to work with collections of data compared to arrays, which have a fixed size. A slice does not store data itself; it only describes a portion of an array. Any modification to a slice affects the original array.

#### **Declaring and Initializing Slices**

Slices can be declared and initialized similarly to arrays, but without specifying a fixed size:

```go
package main

import (
  "fmt"
  "reflect"
)

func main() {
  array1 := [3]int{1, 2, 3}  // Fixed-size array
  slice1 := []int{1, 2, 3}   // Dynamic-size slice

  fmt.Println("Array:", array1)
  fmt.Println("Slice:", slice1)
  fmt.Println("Array Type:", reflect.TypeOf(array1))
  fmt.Println("Slice Type:", reflect.TypeOf(slice1))
}
```

#### **Creating a Slice from an Array**

Slices can be derived from arrays using a range of indices:

```go
package main

import "fmt"

func main() {
  array2 := [5]int{1, 2, 3, 4, 5} // Array
  slice2 := array2[1:3]           // Slice with elements from index 1 to 2

  fmt.Println("Array:", array2)
  fmt.Println("Slice2:", slice2) // Output: [2, 3]
}
```

#### **Slicing with Different Ranges**

Slices can be created using different index ranges:

```go
package main

import "fmt"

func main() {
  array := [5]int{1, 2, 3, 4, 5}
  sliceA := array[:2]  // Elements from index 0 to 1
  sliceB := array[2:]  // Elements from index 2 to end
  sliceC := sliceA[:1] // Extracting a slice from another slice

  fmt.Println("Array:", array)
  fmt.Println("SliceA:", sliceA) // Output: [1, 2]
  fmt.Println("SliceB:", sliceB) // Output: [3, 4, 5]
  fmt.Println("SliceC:", sliceC) // Output: [1]
}
```

#### **Modifying a Slice Affects the Original Array**

Since slices reference an underlying array, modifying a slice also modifies the original array:

```go
package main

import "fmt"

func main() {
  array := [5]int{1, 2, 3, 4, 5}
  slice := array[1:4]

  fmt.Println("Before modification:", array)
  slice[0] = 99 // Modifies array[1]
  fmt.Println("After modification:", array)
  fmt.Println("Slice:", slice) // Output: [99, 3, 4]
}
```

#### **Iterating Over a Slice**

A `for` loop can be used to iterate over a slice:

```go
package main

import "fmt"

func main() {
  slice := []int{10, 20, 30, 40, 50}

  for i, value := range slice {
    fmt.Printf("Index %d: Value %d\n", i, value)
  }
}
```

Slices provide a powerful and efficient way to handle collections of data in Go while maintaining a reference to an underlying array.

#### **Creating Slices Using `make`**

The `make` function allows you to create slices with a predefined length and optional capacity. Unlike declaring slices with literals, `make` provides more control over the internal array backing the slice.

```go
package main

import "fmt"

func main() {
    s := make([]int, 10) // Creates a slice with length 10
    fmt.Println("Slice:", s)

    s[9] = 50
    fmt.Println("Modified Slice:", s)

    s = make([]int, 10, 20) // Creates a slice of length 10, but capacity 20
    fmt.Println("Slice:", s, "Length:", len(s), "Capacity:", cap(s))
}
```

**Explanation:**
- `make([]int, 10)` creates a slice with 10 elements, all initialized to zero.
- The value at index 9 is modified to `50`, showing that slices can be manipulated like arrays.
- `make([]int, 10, 20)` creates a slice of length `10`, but the underlying array has a capacity of `20`.

#### **Appending Elements to a Slice**

The `append` function dynamically adds elements to a slice, automatically resizing it if needed.

```go
package main

import "fmt"

func main() {
    s := make([]int, 10, 20)
    s = append(s, 1, 2, 3, 4, 5, 6, 7, 8, 9, 0)
    fmt.Println("After append:", s, "Length:", len(s), "Capacity:", cap(s))

    s = append(s, 1) // Exceeds initial capacity, Go allocates a larger array
    fmt.Println("After exceeding capacity:", s, "Length:", len(s), "Capacity:", cap(s))
}
```

**Explanation:**
- `append(s, 1, 2, 3, ...)` adds elements to `s`, increasing its length.
- The first `append` call does not exceed capacity (`20`), so the same underlying array is used.
- The second `append` call adds another element beyond capacity, forcing Go to allocate a new array with increased size.

Using `make` and `append` together allows efficient slice handling, balancing pre-allocation and dynamic growth.

### 3.4 - Map

In Go, a `map` is a data structure that associates keys with values. It is very useful for storing and accessing data efficiently, as long as you know the corresponding key. Keys in a `map` must be of a type that is comparable, such as numbers, strings, or other types that support comparison.

#### **Creating and Initializing a Map**

You can create a `map` using the `make` function or initialize it directly with values:

```go
package main

import "fmt"

func main() {
  // Creating an empty map using make
  countries := make(map[string]string)

  // Adding elements to the map
  countries["US"] = "United States"
  countries["BR"] = "Brazil"
  countries["JP"] = "Japan"

  fmt.Println("Countries:", countries)
}
```

#### **Iterating Over a Map**

You can use a `for range` loop to iterate over the keys and values of a `map`:

```go
package main

import "fmt"

func main() {
  countries := map[string]string{
    "US": "United States",
    "BR": "Brazil",
    "JP": "Japan",
  }

  for code, name := range countries {
    fmt.Printf("Country code %s represents %s.\n", code, name)
  }
}
```

#### **Deleting Elements from a Map**

To remove an element from a `map`, use the `delete` function:

```go
package main

import "fmt"

func main() {
  countries := map[string]string{
    "US": "United States",
    "BR": "Brazil",
    "JP": "Japan",
  }

  fmt.Println("Before deletion:", countries)

  // Removing the element with the key "BR"
  delete(countries, "BR")

  fmt.Println("After deletion:", countries)
}
```

#### **Nested Map**

You can create nested maps, where the values of a `map` are other `map` structures:

```go
package main

import "fmt"

func main() {
  citiesByCountry := map[string]map[string]int{
    "US": {
      "New York": 8419600,
      "Los Angeles": 3980400,
    },
    "JP": {
      "Tokyo": 13929286,
      "Osaka": 2715000,
    },
  }

  for country, cities := range citiesByCountry {
    fmt.Printf("Cities in %s:\n", country)
    for city, population := range cities {
      fmt.Printf("  %s has a population of %d.\n", city, population)
    }
  }
}
```

#### **Checking for Key Existence**

You can check if a key exists in a `map` using the second return value when accessing the `map`:

```go
package main

import "fmt"

func main() {
  countries := map[string]string{
    "US": "United States",
    "BR": "Brazil",
  }

  if name, exists := countries["JP"]; exists {
    fmt.Printf("JP represents %s.\n", name)
  } else {
    fmt.Println("JP is not in the map.")
  }
}
```

Maps are a powerful tool in Go for associating data efficiently and flexibly. However, remember that they do not guarantee the order of elements when iterating over them.

## 4 - Functions
