// generics/standard-library/begin/main.go
package main

import (
	"fmt"

	"golang.org/x/exp/constraints"
)

// create a numeric interface with a type set
type numeric interface {
	constraints.Integer | constraints.Float
	grow()
}

// update sum function to use a numeric interface with a type set
func sum[T numeric](a, b T) T {
	return a + b
}

type specialInt int

func (s specialInt) grow() {}

// equal returns true if a and b are equal.
func equal[T comparable](a, b T) bool {
	return a == b
}

func main() {
	// invoke equal with comparable types
	fmt.Println("equal(1, 1):", equal(1, 1))
	fmt.Println("equal(\"one\", \"two\"):", equal("one", "two"))

	// invoke equal with a custom type
	type c struct{ f string }
	fmt.Println("equal(c{f: \"a\"}, c{f: \"a\"}):", equal(c{f: "a"}, c{f: "a"}))
	fmt.Println("equal(c{f: \"a\"}, c{f: \"b\"}):", equal(c{f: "a"}, c{f: "b"}))
}
