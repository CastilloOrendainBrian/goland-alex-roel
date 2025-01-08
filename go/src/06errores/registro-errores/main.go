package main

import (
	"log"
	"os"
)

func main() {
	// log.Fatal("¡Oye, soy un registro de errores")
	// log.SetPrefix("main(): ")
	// log.Panic("¡Oye, soy un registro de errores")
	// log.Print("Oye, soy un log")
	// log.Fatal("¡Oye, soy un registro de errores")
	// log.Print("¿Puedes verme?")
	log.SetPrefix("main(): ")
	file, err := os.OpenFile("info.log", os.O_CREATE|os.O_APPEND|os.O_WRONLY, 0644)
	if err != nil {
		log.Fatal(err)
	}

	defer file.Close()

	log.SetOutput(file)
	log.Print("¡Hola, soy un log!")
}
