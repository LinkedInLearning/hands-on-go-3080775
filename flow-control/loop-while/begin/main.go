// flow-control/loop-while/begin/main.go
package main

import "fmt"

func main() {
	// initialize a variable to check if the loop should continue
	n := 5

	// iterate while the condition is true
	for n > 0 {
		fmt.Println(n)
		n--
	}
}
