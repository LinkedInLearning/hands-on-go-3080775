// flow-control/loop-basic/begin/main.go
package main

import "fmt"

func main() {
	// declare a string to iterate over
	//
	greet := "Hello"

	// iterate over the string with basic for loop
	//
	for i := 0; i < len(greet); i++ {
		fmt.Println(i, ":", string(greet[i]))
	}
}
