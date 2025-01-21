package main

import "fmt"

// Función que retorna una función
func incrementar() func() int {
	i := 0
	return func() int {
		i++
		return i
	}
}

func main() {
	nextInt := incrementar()

	fmt.Println(nextInt()) // 1
	fmt.Println(nextInt()) // 2
	fmt.Println(nextInt()) // 3
	fmt.Println(nextInt()) // 4
	fmt.Println(nextInt()) // 5
}
