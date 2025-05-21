package main

import "fmt"

type Item struct {
	productId int
	quantity  int
	price     float64
}

type Order struct {
	userId int
	items  []Item
}

func (o Order) TotalPrice() float64 {
  total := 0.0

  for _, item := range o.items {
    total += item.price * float64(item.quantity)
  }

  return total
}

func main() {
  order := Order{
    userId: 1,
    items: []Item{
      {productId: 101, quantity: 2, price: 50.0},
      {productId: 102, quantity: 1, price: 100.0},
    },
  }

  fmt.Println(order.TotalPrice())
}
