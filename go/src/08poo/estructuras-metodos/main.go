package main

import "library/book"

func main() {
	// var myBook = book.Book{
	// 	Title:  "The Art of Computer Programming",
	// 	Author: "Donald Knuth",
	// 	Pages:  700,
	// }
	myBook := book.NewBook("The Art of Computer Programming", "Donald Knuth", 700)

	myBook.PrintInfo()
}
