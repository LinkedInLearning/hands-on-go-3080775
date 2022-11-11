// functions/begin/main.go
package main

import (
	"errors"
	"fmt"
	"strconv"
)

// simple greet function
func greet() string {
	return "Greet funtion called"
}

//

// greetWithName returns a greeting with the name
func greetWithName(name string) string {
	return ("Hola " + name)
}

//

// greetWithName returns a greeting with the name and age of the person
func greetWithNameAge(name string, age int) (greeting string) {
	//return ("Hola " + name + ". Your age is " + strconv.Itoa(age))
	greeting = "Hola " + name + ". Your age is " + strconv.Itoa(age)
	return
}

// divide divides two numbers and returns the result
// if the second number is zero, it returns  error
func divide(a int, b int) (int, error) {
	if b == 0 {
		return 0, errors.New("divide by zero not allowed. you dont know maths")
	}

	return a / b, nil

}

//

func main() {
	// invoke greet function
	fmt.Println(greet())

	// invoke greetWithName function
	fmt.Println(greetWithName("mi amigo"))

	fmt.Println(greetWithNameAge("Yoyo", 23))
	// invoke divide function
	fmt.Println(divide(10, 2))

	// invoke divide function with zero denominator to get an error
	fmt.Println(divide(5, 0))
}
