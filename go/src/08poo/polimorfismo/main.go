package main

import (
	"library/book"
)

func main() {
	var myBook = book.NewBook("The Art of Computer Programming", "Donald Knuth", 700)

	var myTextBook = book.NewTextBook("The C Programming Language", "Brian Kernighan", 300, "Prentice Hall", "Intermediate")

	// myBook.PrintInfo()
	// myTextBook.PrintInfo()

	book.Print(myBook)
	book.Print(myTextBook)
}
