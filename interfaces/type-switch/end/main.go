package main

import "fmt"

func whatAmI(val interface{}) string {
	switch val.(type) {
	case int:
		return fmt.Sprintf("%v is an int", val)
	case string:
		return fmt.Sprintf("%v is a string", val)
	default:
		return fmt.Sprintf("%v is of an unsupported type", val)
	}
}

func main() {
	// invoke whatAmI function
	fmt.Println(whatAmI(1))
	fmt.Println(whatAmI("hello"))
	fmt.Println(whatAmI(true))
}
