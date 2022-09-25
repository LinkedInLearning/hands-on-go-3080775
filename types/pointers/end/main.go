// basic-types/pointers/end/main.go
package main

import "fmt"

func main() {
	// create a variable of type *T where T is an int
	var a *int

	// declare and assign `b` variable of type int
	b := 100

	// assign the address of b to a
	a = &b

	// print out the value of a which is the address of b
	fmt.Println(a)

	// print out the value at the address of b
	fmt.Println(*a)
}
