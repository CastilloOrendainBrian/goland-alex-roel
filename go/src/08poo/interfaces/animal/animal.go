package animal

import "fmt"

type Animal interface {
	Sonido()
}

// Estructura de perro y sus métodos
type Perro struct {
	Nombre string
}

func (p *Perro) Sonido() {
	fmt.Println(p.Nombre + "Guau guau")
}

// Estructura de gato y sus métodos
type Gato struct {
	Nombre string
}

func (g *Gato) Sonido() {
	fmt.Println(g.Nombre + "Miau miau")
}

// Función para imprimir el sonido de un animal
func HacerSonido(a Animal) {
	a.Sonido()
}
