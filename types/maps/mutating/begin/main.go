// types/maps/mutating/begin/main.go
package main

import "fmt"

type author struct {
	name string
}

func main() {
	var authors = map[string]author{
		"tm": {name: "Toni Morrison"},
		"ma": {name: "Marcus Aurelius"},
	}
	fmt.Printf("%v\n", authors)

	// change the value of an author in the map
	authors["tm"] = author{name: "James Baldwind"}

	// print the map
	fmt.Printf("%v\n", authors)

	// delete a key from the map
	delete(authors, "ma")

	// print the map
	fmt.Printf("%v\n", authors)
}
