package main

import "fmt"

type Persona struct {
	nombre  string
	edad    int
	correos string
}

func (p *Persona) diHola() {
	fmt.Println("Hola, mi nombre es", p.nombre)
}

func main() {
	// var x int = 10
	// var p *int = &x
	// fmt.Println(x)
	// fmt.Println(p)
	// editar(&x)
	// fmt.Println(x)
}

func editar(x *int) {
	*x = 20
}

// func main() {
// 	var p Persona
// 	p.nombre = "Juan"
// 	p.edad = 30
// 	p.correos = "juan@mail.com"
// 	p := Persona{"Juan", 30, "juan@mail.com"}
// 	p.nombre = "Pedro"
// 	fmt.Println(p)

// 	p2 := Persona{"Miguel", 23, "miguel@mail.com"}
// 	fmt.Println(p2)

// }

// func main() {
// var a [5]int
// a[2] = 7
// a[4] = 8

// var a = [5]int{1, 2, 3, 4, 5}

// var a = [...]int{1, 2, 3, 4, 5}

// for i := 0; i < len(a); i++ {
// 	fmt.Println(a[i])
// }

// for index, value := range a {
// 	fmt.Printf("Índice: %d, Valor: %d\n", index, value)
// }

// var matriz = [3][3]int{{1, 2, 3}, {4, 5, 6}, {7, 8, 9}}
// fmt.Println(matriz)

// var a []int
// diasSemana := []string{"Lunes", "Martes", "Miércoles", "Jueves", "Viernes", "Sábado", "Domingo"}

// diaRebanada := diasSemana[1:5]
// fmt.Println(diaRebanada)

// diaRebanada = append(diaRebanada, "Día 8", "Día 9", "Día 10")
// fmt.Println(diaRebanada)

// diaRebanada = append(diaRebanada[:2], diaRebanada[3:]...)
// fmt.Println(diaRebanada)

// fmt.Println(len(diaRebanada))
// fmt.Println(cap(diaRebanada))

// nombres := make([]string, 5, 10)
// nombres[0] = "Juan"
// fmt.Println(nombres)

// rebanada1 := []int{1, 2, 3, 4, 5}
// rebanada2 := make([]int, 5)

// copy(rebanada2, rebanada1)

// fmt.Println(rebanada1)
// fmt.Println(rebanada2)

// colors := map[string]string{
// 	"red":   "#ff0000",
// 	"green": "#4bf745",
// 	"blue":  "#0000ff",
// }

// fmt.Println(colors)
// colors["white"] = "#ffffff"
// fmt.Println(colors)

// if valor, ok := colors["rede"]; ok {
// 	fmt.Println(valor)
// } else {
// 	fmt.Println("No existe")
// }

// delete(colors, "red")
// fmt.Println(colors)

// for clave, valor := range colors {
// 	fmt.Printf("Clave: %s, Valor: %s\n", clave, valor)
// }
// }
