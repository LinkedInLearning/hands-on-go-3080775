// types/pointers/test/main.go
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

	d := getNewPointerRefactor(a)
	// print out the value of a which is the address
	fmt.Println(a)
	fmt.Println(d)
	b = 300
	// print out the value at the address
	fmt.Println(*a)
	fmt.Println(*d)
}

func getNewPointer(a *int) *int {
	var b int = *a
	var c *int = &b
	return c
}

func getNewPointerRefactor(a *int) *int {
	b := *a
	return &b
}

func noNewPointer(a *int) *int {
	b := &*a
	return b
}
