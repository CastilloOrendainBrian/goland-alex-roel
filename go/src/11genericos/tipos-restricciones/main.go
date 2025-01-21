package main

import "fmt"

type integer int

func Sum[T ~int | ~float64](nums ...T) T {
	var total T
	for _, v := range nums {
		total += v
	}
	return total
}

func main() {
	var num1, num2 integer = 100, 300
	fmt.Println(Sum(1, 2, 3, 4, 5))
	fmt.Println(Sum(1.1, 2.2, 3.3, 4.4, 5.5))
	fmt.Println(Sum(num1, num2))
}
