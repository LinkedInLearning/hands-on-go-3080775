package main

import "fmt"

// define a struct type for author
type author struct {
	first string
	last  string
}

func main() {
	// intialize author
	a := author{
		first: "James",
		last:  "Baldwin",
	}

	// print the author
	fmt.Printf("%#v\n", a)
}
