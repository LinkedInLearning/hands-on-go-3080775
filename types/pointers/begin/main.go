// types/pointers/begin/main.go
package main

import "fmt"

func main() {
	// create a variable of type *T where T is an
	var a *int

	// declare and assign `b` variable of type int
	b := 100

	// assign the address of b to a
	a = &b

	// print out the value of a which is the address of b
	fmt.Println(a)

	// print out the value at the address of b
	fmt.Println(*a)

	c := noNewPointer(a)
	fmt.Println(c)
	fmt.Println(*c)
	fmt.Println(*a)
	fmt.Println(b)

	//// another test
	d := newPointer(a)
	fmt.Println(d)
	fmt.Println(a)
	b = 300
	fmt.Println(*d)
	fmt.Println(*a)

	// another test
	var myauthor *author = &author{name: "Pepe"}
	jonauthor := newAuthor(myauthor)
	fmt.Printf("address: %p - value: %s\n", myauthor, *myauthor)
	fmt.Printf("address: %p - value: %s\n", jonauthor, *jonauthor)
}

func noNewPointer(a *int) *int {
	b := *&a
	return b
}

func newPointer(a *int) *int {
	var b int = *a
	var c *int = &b
	return c
}

func newAuthor(oldauthor *author) *author {
	var newauthor *author = &author{name: "Jon"}
	return newauthor
}

type author struct{ name string }
