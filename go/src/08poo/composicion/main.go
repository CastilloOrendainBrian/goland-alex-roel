package main

import (
	"fmt"
	"library/book"
)

func main() {
	myBook := book.NewBook("The Art of Computer Programming", "Donald Knuth", 700)
	myBook.PrintInfo()

	myBook.SetTitle("The Go Programming Language")
	fmt.Println("New title:", myBook.GetTitle())

	myTextBook := book.NewTextBook("The C Programming Language", "Brian Kernighan", 300, "Prentice Hall", "Intermediate")

	myTextBook.PrintInfo()
}
