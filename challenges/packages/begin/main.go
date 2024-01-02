package main

import (
	"fmt"

	"github.com/jboursiquot/go-proverbs"
)

func getRandomProverb() string {
	return proverbs.Random().Saying
}

func main() {
	fmt.Print(getRandomProverb())
}
