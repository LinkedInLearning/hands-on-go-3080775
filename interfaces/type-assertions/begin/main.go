// interfaces/type-assertions/begin/main.go
package main

import "fmt"

func main() {
	// perform a type assertion
	var text any = "Hola"
	fmt.Println(text.(string))

	// perform a type assertion and handle the error
	var number any = "hola"

	if _, ok := number.(int); !ok {
		fmt.Println("Number is not a number")
	}
}
