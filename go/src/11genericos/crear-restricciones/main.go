package main

import (
	"fmt"

	"golang.org/x/exp/constraints"
)

type integer int

type Numbers interface {
	~int | ~float64 | ~float32 | ~uint
}

func Sum[T constraints.Integer | constraints.Float](nums ...T) T {
	var total T
	for _, v := range nums {
		total += v
	}
	return total
}

func main() {
	var num1, num2 integer = 100, 300
	fmt.Println(Sum[uint](1, 2, 3, 4, 5))
	fmt.Println(Sum[float32](1.1, 2.2, 3.3, 4.4, 5.5))
	fmt.Println(Sum(num1, num2))
}
