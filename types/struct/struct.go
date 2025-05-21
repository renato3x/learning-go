package main

import "fmt"

/*
Structs are used to create complex data types that group together related data.
Structs are defined using the `type` keyword followed by the name of the struct and the `struct` keyword.
Structs can contain fields of different types, and each field has a name and a type.

Structs is not a class, but it is similar to a class in that it can contain methods.
*/
type Product struct {
	name     string
	price    float64
	discount float64
}

/* 
  This is a method that belongs to the Product struct.
  The structure (p Product) means that this method can only be called on a Product struct.
  It's similar to a method in a class.
  The technical term for this is a "receiver".
  The receiver is the instance of the struct that the method is called on.

  p is a receiver, which is a special variable that refers to the instance of the struct that the method is called on.
  The receiver is defined in parentheses before the method name.
*/
func (p Product) PriceWithDiscount() float64 {
  return p.price - (1 - p.discount)
}

func main() {
  product := Product{
    name:     "Pencil",
    price:    100.0,
    discount: 0.1,
  }

  fmt.Println(product.price, product.PriceWithDiscount())

  var product2 Product
  product2 = Product{
    name:     "Laptop",
    price:    1000.0,
    discount: 0.2,
  }
  fmt.Println(product2.price, product2.PriceWithDiscount())
}
