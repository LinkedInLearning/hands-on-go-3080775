// challenges/types-composite/begin/main.go
package main

import (
	"fmt"

	"github.com/davecgh/go-spew/spew"
)

// define an author type with a name
//
type author struct {
	name string
}

// define a book type with a title and author
//
type book struct {
	title string
	author
}

// define a library type to track a list of books
//
type library map[string][]book

// define addBook to add a book to the library
//
func (l library) addBook(b book) {
	l[b.author.name] = append(l[b.author.name], b)
}

// define a lookupByAuthor function to find books by an author's name
//
func (l library) lookupByAuthor(a author) []book {
	return l[a.name]
}

func main() {
	// create a new library
	lib := library{}

	// add 2 books to the library by the same author

	tolkien := author{name: "Tolkien"}
	lib.addBook(book{title: "Lord of the rings", author: tolkien})
	lib.addBook(book{title: "Silmarillon", author: tolkien})

	// add 1 book to the library by a different author
	martin := author{name: "George R. Martin"}
	lib.addBook(book{title: "Game of thrones", author: martin})

	// dump the library with spew
	//
	spew.Dump(lib)

	// lookup books by known author in the library
	//
	books := lib.lookupByAuthor(author{name: tolkien.name})
	spew.Dump(books)
	
	// print out the first book's title and its author's name
	//
	if len(books) != 0 {
		fmt.Printf("Book name :" + books[0].name + ", author's name: " + books[0].author.name)
	}
}
