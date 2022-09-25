package main

import "fmt"

type author struct {
	name string
}

func main() {
	// declare a map of string keys and author values
	var authors map[string]author

	// initialize the map with make
	authors = make(map[string]author)

	// add authors to the map
	authors["tm"] = author{name: "Toni Morrison"}
	authors["ma"] = author{name: "Marcus Aurelius"}

	// alternative syntax to declare and initialize the author map
	// authors := map[string]author{
	// 	"tm": author{name: "Toni Morrison"},
	// 	"ma": author{name: "Marcus Aurelius"},
	// }

	// print the map with fmt.Printf
	fmt.Printf("%#v\n", authors)

	// read a value from the map with a known key
	fmt.Println(authors["tm"])
}
