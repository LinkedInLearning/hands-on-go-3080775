// packages/basics/main.go
package main

import (
	"fmt"
	"strconv"
	"time"
)

func greet() string {
	return "Hello!"
}

func greetWithName(name string) string {
	return "Hello, " + name + "!"
}

func greetWithNameAndAge(name string, age int) (greeting string) {
	greeting = "Hello, my name is " + name + " and I'm " + strconv.Itoa(age) + " years old"
	return
}
func main() {
	//fmt.Println("Hello Gopher!")
	//fmt.Println(greet())
	//fmt.Println(greetWithName("Jon Cabrera"))
	fmt.Println(greetWithNameAndAge("Jesus", 33))
	fmt.Printf("Today is %s", time.Now().Weekday())
}
