// packages/imports/begin/main.go
package main

// import the fmt package from the standard library
import (
	"errors"
	"fmt"
	"strconv"
)

func greet() string {
	return "good day sir"
}

func greetWithName(name string) string {
	return "Hello, " + name + "!"
}

// named return parameter
func greetWithNameAndAge(name string, age int) (greeting string) {
	greeting = "Hello, my name is " + name + " and I am " + strconv.Itoa(age) + " years old."
	return //naked return
}

func divide(a, b int) (int, error) {
	if b == 0 {
		return 0, errors.New("cannot divide by 0")
	}
	return a / b, nil
}

// import the time package from the standard library

func main() {
	// fmt.Println(greet())
	// fmt.Println(greetWithName("Ricky"))
	// fmt.Println(greetWithNameAndAge("Ricky", 17))
	fmt.Println(divide(10, 2))
	fmt.Println(divide(5, 0))

	// use the fmt package to print the string "Hello Gopher!"

	// use the time package to print the current weekday
}
