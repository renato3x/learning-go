package main

import (
	"encoding/json"
	"fmt"
)

type Product struct {
	Id    int      `json:"id"` // this a mapping of JSON key to struct field. This isn't mandatory, but it's a good practice
	Name  string   `json:"name"`
	Price int      `json:"price"`
	Tags  []string `json:"tags"`
}

func main() {
	// struct to JSON
	p1 := Product{
		Id:    1,
		Name:  "Laptop",
		Price: 100,
		Tags:  []string{"Promotion", "Electronics"},
	}
	p1Json, _ := json.Marshal(p1)
	fmt.Println(string(p1Json))

	jsonString := `{"id":2,"name":"Phone","price":200,"tags":["Promotion","Electronics"]}`
	json.Unmarshal([]byte(jsonString), &p1)

	fmt.Println(p1)
}
