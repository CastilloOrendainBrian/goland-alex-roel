package main

import (
	"fmt"
	"math"
)

func main() {
	var lado1, lado2 float64
	const precision = 2

	//Entrada de datos
	fmt.Print("Ingrese el valor del lado 1: ")
	fmt.Scan(&lado1)
	fmt.Print("Ingrese el valor del lado 2: ")
	fmt.Scan(&lado2)

	//Proceso
	area := (lado1 * lado2) / 2
	hipotenusa := math.Sqrt(math.Pow(lado1, 2) + math.Pow(lado2, 2))
	perimetro := lado1 + lado2 + hipotenusa

	//Salida
	// fmt.Printf("El area del triangulo es: %.2f\n", area)
	// fmt.Printf("El perimetro del triangulo es: %.2f\n", perimetro)
	// fmt.Printf("La hipotenusa del triangulo es: %.2f\n", hipotenusa)

	fmt.Printf("El area del triangulo es: %.*f\n", precision, area)
	fmt.Printf("El perimetro del triangulo es: %.*f\n", precision, perimetro)
	fmt.Printf("La hipotenusa del triangulo es: %.*f\n", precision, hipotenusa)
}
