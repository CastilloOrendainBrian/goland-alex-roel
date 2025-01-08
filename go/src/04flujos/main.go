package main

func main() {
	// t := time.Now()
	// hora := t.Hour()

	// if hora < 12 {
	// 	fmt.Println("Bom dia!")
	// } else if hora < 18 {
	// 	fmt.Println("Boa tarde!")
	// } else {
	// 	fmt.Println("Boa noite!")
	// }

	// if t := time.Now(); t.Hour() < 12 {
	// 	fmt.Println("Bom dia!")
	// } else if t.Hour() < 18 {
	// 	fmt.Println("Boa tarde!")
	// } else {
	// 	fmt.Println("Boa noite!")
	// }

	// switch t := time.Now(); {
	// case t.Hour() < 12:
	// 	fmt.Println("Bom dia!")
	// case t.Hour() < 18:
	// 	fmt.Println("Boa tarde!")
	// default:
	// 	fmt.Println("Boa noite!")
	// }

	// os := runtime.GOOS
	// switch os := runtime.GOOS; os {
	// case "darwin":
	// 	fmt.Println("Mac OS")
	// case "linux":
	// 	fmt.Println("Linux")
	// default:
	// 	fmt.Println("Windows")
	// }

	// var i int
	// for i := 1; i <= 10; i++ {
	// 	if i == 5 {
	// 		continue
	// 	}
	// 	fmt.Println(i)
	// }
	saludo := hello("Alex")
	println(saludo)
	sum, mul := calc(2, 3)
	println("La suma es: ", sum)
	println("La multiplicación es: ", mul)
}

func hello(name string) string {
	return "Hello, " + name
}

func calc(a, b int) (sum, mul int) {
	sum = a + b
	mul = a * b

	return
}
