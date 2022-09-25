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
	authors["tm"] = author{name: "James Baldwin"}

	// print the map
	fmt.Printf("%v\n", authors)

	// delete the author with the key "ma"
	delete(authors, "tm")

	// print the map
	fmt.Printf("%v\n", authors)
}
