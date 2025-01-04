package main

import (
	"fmt"
	"math"
)

func tipos() {
	// fmt.Println(math.MaxUint8)
	// fullName := "Alex Roel \t(alias \"roelcode\")\n"
	// fmt.Println(fullName)

	// var a byte = 'A'
	// fmt.Println(a)

	// s := "Hello, World!"
	// fmt.Println(s[0])

	// var r rune = '♥'
	// fmt.Println(r)

	// var (
	// 	defaultInt     int
	// 	defaultUint    uint
	// 	defaultFloat   float32
	// 	defaultComplex bool
	// 	defaultString  string
	// )

	// fmt.Println(defaultInt, defaultUint, defaultFloat, defaultComplex, defaultString)

	// var integer16 int16 = 50
	// var integer32 int32 = 100

	// fmt.Println(int32(integer16) + integer32)

	// s := "100"
	// i, _ := strconv.Atoi(s)

	// fmt.Println(i)

	// n := 42
	// s = strconv.Itoa(n)

	// fmt.Println(s + s)

	// name := "Alex"
	// age := 25

	// fmt.Printf("Hola, me llamo %s y tengo %d años.\n", name, age)

	// greeting := fmt.Sprintf("Hola, me llamo %s y tengo %d años.", name, age)

	// fmt.Println(greeting)

	var name string

	var age int

	fmt.Print("ingrese su nombre: ")
	fmt.Scanln(&name) // recibe la referencia en memoria de una variable

	fmt.Print("ingrese su edad: ")
	fmt.Scanln(&age)

	greeting := fmt.Sprintf("Hola, me llamo %s y tengo %d años.", name, age)
	fmt.Println(greeting)

	fmt.Println(math.Pi)
}
