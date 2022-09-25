// testing/basics/begin/main.go
package main

import "golang.org/x/exp/constraints"

type numeric interface {
	constraints.Integer | constraints.Float
}

// generic sum sums a slice of numbers
func sum[T numeric](numbers ...T) T {
	var s T
	for _, n := range numbers {
		s += n
	}
	return s
}
