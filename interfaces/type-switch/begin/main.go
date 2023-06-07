// interfaces/type-switch/begin/main.go
package main

import "fmt"

// define whatAmI which takes in an argument of any type and returns inforamtion about the underlying value's type
func whatAmI(anyType any) string {
	switch anyType.(type) {
	case int:
		return fmt.Sprintf("The type is int")
	case string:
		return fmt.Sprintf("The type is string")
	default:
		return fmt.Sprintf("The type is other")
	}
}

func main() {
	// invoke whatAmI function
	fmt.Println(whatAmI(1))
	fmt.Println(whatAmI("hello"))
	fmt.Println(whatAmI(true))
}
