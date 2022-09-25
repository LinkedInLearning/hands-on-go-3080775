package main

import "fmt"

// create a numeric interface with a type set
type numeric interface {
	~int | ~float64
	grow()
}

// update sum function to use a numeric interface with a type set
func sum[T numeric](a, b T) T {
	return a + b
}

type specialInt int

func (s specialInt) grow() {}

func main() {
	one := specialInt(1)
	two := specialInt(2)
	fmt.Println(sum(one, two))
}
