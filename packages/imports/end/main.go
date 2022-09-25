package main

import (
	// import the fmt package from the standard library
	"fmt"

	// import the time package from the standard library
	"time"
)

func main() {
	// use the fmt package to print the string "Hello Gopher!"
	fmt.Println("Hello Gopher!")

	// use the time package to print the current weekday
	fmt.Printf("Today is %s\n", time.Now().Weekday())
}
