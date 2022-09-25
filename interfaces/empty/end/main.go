package main

import "fmt"

// define whatAmI which takes in an argument of any type and returns its underlying value's type
func whatAmI(val interface{}) string {
	return fmt.Sprintf("%v, %T", val, val)
}

func main() {
	// invoke whatAmI with an int
	fmt.Println(whatAmI(1))

	// invoke whatAmI with a string
	fmt.Println(whatAmI("hello"))
}
