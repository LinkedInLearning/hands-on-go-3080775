// generics/type-parameters/begin/main.go
package main

import "fmt"

func sumInts(a, b int) int {
	return a + b
}

func sumFloats(a, b float64) float64 {
	return a + b
}

// create generic sum function with type parameter T constrained to int and float64 types
func sum[T int | float64](a, b T) T {
	return a + b
}

func sumAlias[T ~int | ~float64](a, b T) T {
	return a + b
}

type intalias int

func main() {
	// non-generic sum int function
	fmt.Println(sumInts(1, 2))

	// non-generic sum float function
	fmt.Println(sumFloats(1.3, 2.2))

	// call on generic sum function
	fmt.Println(sum[int](1, 2))
	fmt.Println(sum[float64](3.1, 2.4))

	// define a compatible custom type call on generic sum function with it
	//
	aliasf := intalias(5)
	aliass := intalias(10)
	fmt.Println(sumAlias(aliasf, aliass))
}

// list is a singly-linked list that holds values of any type
type list[T any] struct {
	next *list[T]
	elem T
}