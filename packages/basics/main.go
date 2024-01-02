// packages/basics/main.go
package main

func greet() string {
	return "good day sir"
}

import (
	"fmt"
)

func main() {

	
	fmt.Println("hello Louis")
	fmt.Println("Today is %s", time.Now().Weekday())
}
