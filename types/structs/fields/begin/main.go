// types/structs/fields/begin/main.go
package main

import "fmt"

// define a struct type for author
type author struct {
	first, last string
}

func main() {
	// intialize author
	a := author{
		first: "James",
		last:  "Beckett",
	}

	// print the author
	fmt.Printf("%#v\n", a)
}
