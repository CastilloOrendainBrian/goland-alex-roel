package main

import (
	"fmt"
	"math/rand"
)

func main() {
	// rand.Seed(time.Now().UnixNano())
	// fmt.Println(rand.Intn(100))
	jugar()
}

func jugar() {
	numAleatorio := rand.Intn(100)
	var numIngresado int
	var intentos int
	const maxIntentos = 10

	for intentos < maxIntentos {
		intentos++
		fmt.Printf("Intento %d: Ingrese un número: ", intentos)
		fmt.Scan(&numIngresado)

		if numIngresado == numAleatorio {
			fmt.Println("¡Felicidades! Adivinaste el número.")
			jugarNuevamente()
			break
		} else if numIngresado < numAleatorio {
			fmt.Println("El número ingresado es menor al número aleatorio.")
		} else {
			fmt.Println("El número ingresado es mayor al número aleatorio.")
		}
	}

	fmt.Printf("El número aleatorio era: %d\n", numAleatorio)
	jugarNuevamente()
}

func jugarNuevamente() {
	var respuesta string
	fmt.Print("¿Desea jugar nuevamente? (s/n): ")
	fmt.Scan(&respuesta)

	if respuesta == "s" {
		jugar()
	} else if respuesta == "n" {
		fmt.Println("Gracias por jugar.")
	} else {
		fmt.Println("Respuesta inválida.")
		jugarNuevamente()
	}
}
