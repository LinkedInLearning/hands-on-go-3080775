// generics/type-sets/begin/main.go
package main

import "fmt"

// create a numeric interface with a type set

// update sum function to use a numeric interface with a type set
func sum[T ~int | ~float64](a, b T) T {
	return a + b
}

type specialInt int

func main() {
	one := specialInt(1)
	two := specialInt(2)
	fmt.Println(sum(one, two))
}
