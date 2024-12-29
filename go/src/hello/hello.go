package main

import (
	"fmt"

	"rsc.io/quote"
)

// Declare a constant

const Pi float32 = 3.14

const (
	X = 100
	Y = 0b1010 // Binary
	Z = 0o12   // Octal
	W = 0xFF   // Hexadecimal
)

const (
	Domingo = iota + 1
	Lunes
	Martes
	Miercoles
	Jueves
	Viernes
	Sabado
)

func main() {
	fmt.Println("Hello, World!")
	fmt.Println(quote.Go())

	// Declare a variable

	// var firstName, lastName string
	// var age int

	// var (
	// 	firstName, lastName string
	// 	age                 int
	// )

	// var (
	// 	firstName = "John"
	// 	lastName  = "Doe"
	// 	age       = 25
	// )

	// var firstName, lastName, age = "John", "Doe", 25

	firstName, lastName, age := "John", "Doe", 25

	// firstName = "John"
	// lastName = "Doe"
	// age = 25

	fmt.Println(firstName, lastName, age)
	fmt.Println(Pi)
	fmt.Println(X, Y, Z, W)
	fmt.Println(Domingo, Lunes, Martes, Miercoles, Jueves, Viernes, Sabado)
}
