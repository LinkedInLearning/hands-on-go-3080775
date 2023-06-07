// types/maps/initialization/begin/main.go
package main

import "fmt"

type author struct {
	name string
}

func main() {
	// declare a map of string keys and author values
	//
	var authors map[string]author

	// initialize the map with make
	//
	authors = make(map[string]author)

	// add authors to the map
	//
	authors["1"] = author{name: "Jimmy"}
	authors["2"] = author{name: "Alfred"}

	// print the map with fmt.Printf
	//
	fmt.Println(authors)

	// read a value from the map with a known key
	fmt.Println(authors["1"])
}
