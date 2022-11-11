// packages/imports/begin/main.go
package main

import (
	"fmt"
	"time"
)

// import the fmt package from the standard library

// import the time package from the standard library

func main() {
	// use the fmt package to print the string "Hello Gopher!"
	fmt.Println("Hello from fmt package import")
	fmt.Print(time.Now().Weekday())
	// use the time package to print the current weekday
}
