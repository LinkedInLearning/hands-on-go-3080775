package main

import (
	"fmt"

	"github.com/jboursiquot/go-proverbs"
)

func main() {
	p, err := proverbs.Nth(4)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Printf("%+v\n", p)
}
