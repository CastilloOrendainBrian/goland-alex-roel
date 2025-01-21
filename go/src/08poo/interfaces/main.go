package main

import "library/animal"

func main() {
	// miPerro := animal.Perro{Nombre: "Firulais "}
	// miGato := animal.Gato{Nombre: "Garfield "}

	// animal.HacerSonido(&miPerro)
	// animal.HacerSonido(&miGato)

	animales := []animal.Animal{
		&animal.Perro{Nombre: "Firulais "},
		&animal.Gato{Nombre: "Garfield "},
		&animal.Perro{Nombre: "Rex "},
		&animal.Gato{Nombre: "Tom "},
	}

	for _, animal := range animales {
		animal.Sonido()
	}
}
