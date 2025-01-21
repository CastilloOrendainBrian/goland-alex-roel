package main

import "fmt"

func PrintList(list ...any) {
	for _, v := range list {
		fmt.Println(v)
	}
}

func main() {
	PrintList("Alexys", "Leon", "Vargas")
	PrintList(1, 2, 3, 4, 5)
	PrintList(1, "Hello", 3.14, true)
}
