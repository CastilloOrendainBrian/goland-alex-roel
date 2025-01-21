package main

import "fmt"

// Función Variádica
func suma(name string, nums ...int) int {
	total := 0
	for _, num := range nums {
		total += num
	}

	fmt.Printf("Hola %s, la suma es: %d\n", name, total)

	return total
}

func imprimirDatos(datos ...interface{}) {
	for _, dato := range datos {
		fmt.Printf("%T - %v\n", dato, dato)
	}
}

func main() {
	fmt.Println(suma("Juan", 1, 2, 3, 4, 5))
	fmt.Println(suma("Pedro", 1, 2, 3, 4, 5, 6, 7, 8, 9, 10))
	imprimirDatos(1, "Hola", 3.14, true)
}
