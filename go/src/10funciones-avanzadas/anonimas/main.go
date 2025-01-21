package main

import "fmt"

// func saludar(name string, f func(string)) {
// 	f(name)
// }

func duplicar(n int) int {
	return n * 2
}

func triplicar(n int) int {
	return n * 3
}

func aplicar(f func(int) int, n int) int {
	return f(n)
}

func main() {
	// Declaración de una función anónima
	// func() {
	// 	println("Hola mundo")
	// }()

	// Declaración de una función anónima con parámetros
	// func(s string) {
	// 	println(s)
	// }("Hola mundo")

	// saludo := func(s string) {
	// 	fmt.Printf("Hola %s\n", s)
	// }

	// saludar("Mundo", saludo)

	r1 := aplicar(duplicar, 5)
	r2 := aplicar(triplicar, 5)
	fmt.Println(r1, r2)
}
