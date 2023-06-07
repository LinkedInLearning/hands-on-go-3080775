// types/variables/begin/main.go
package main

import "fmt"

// declare package-level variables of type int
//
var d, e, f int

// declare package-level variables of type bool and override the default values (also known as "zero")
//
var x, y, z = true, false, true

func main() {
	// print ints
	//
	fmt.Print("d, e, f:", d, e, f)

	// override the default value of a package-level variable
	d = 1_000_000
	fmt.Printf("d: %d\n", d)

	// declare and initialize variables of similar names as booleans but of local scope
	x, y, z := 0, 1.25, "hello"
	fmt.Println("x, y, z:", x, y, z)

	// print the package-level variables x, y, and z
	printPackageLevelVariables()
}

func printPackageLevelVariables() {
	fmt.Println("x, y, z:", x, y, z)
}
