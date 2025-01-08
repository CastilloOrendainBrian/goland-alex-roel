package main

import (
	"fmt"
	"os"
)

func divide(dividendo, divisor int) (int, error) {
	if divisor == 0 {
		// return 0, errors.New("No se puede dividir por cero")
		return 0, fmt.Errorf("No se puede dividir por cero")
	}
	return dividendo / divisor, nil
}

func main() {
	// str := "123f"
	// num, err := strconv.Atoi(str)
	// if err != nil {
	// 	fmt.Println("Error:", err)
	// 	return
	// }

	// fmt.Println("Número:", num)

	// result, err := divide(18, 0)
	// if err != nil {
	// 	fmt.Println("Error:", err)
	// 	return
	// }

	// fmt.Println("Resultado:", result)

	// defer fmt.Println(3)
	// defer fmt.Println(2)
	// defer fmt.Println(1)

	file, err := os.Create("hola.txt")
	if err != nil {
		fmt.Println(err)
		return
	}

	defer file.Close()

	_, err = file.Write([]byte("Hola, Brian"))
	if err != nil {
		fmt.Println(err)
		// file.Close()
		return
	}
	// file.Close()
}
