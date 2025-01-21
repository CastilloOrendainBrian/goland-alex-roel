package main

import (
	"fmt"
)

type Product[T uint | string] struct {
	Id    T
	Desc  string
	Price float32
}

func main() {
	product1 := Product[uint]{Id: 1, Desc: "Product 1", Price: 12.34}
	product2 := Product[string]{Id: "A", Desc: "Product 2", Price: 56.78}
	fmt.Println(product1)
	fmt.Println(product2)
}
