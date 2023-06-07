// challenges/flow-control/begin/main.go
package main

import (
	"fmt"
	"os"
	"unicode"
	"github.com/davecgh/go-spew/spew"
	"strings"
)

func main() {
	// handle any panics that might occur anywhere in this function
	//
	defer func() {
		if err := recover(); err != nil {
			fmt.Println("Something wrong happened, but don't panic", err)
		}
	}()

	// use os package to read the file specified as a command line argument
	fileName := os.Args[1]
	file, err := os.ReadFile(fileName)
	if err != nil {
		panic(fmt.Errorf("failed to read file: %s", err))
	}
	
	// convert the bytes to a string
	text := string(file)

	// initialize a map to store the counts
	count := make(map[string]int)

	// use the standard library utility package that can help us split the string into words
	words := strings.Fields(text)

	// capture the length of the words slice
	//
	count["words"] = len(words)

	// use for-range to iterate over the words slice and for each word, count the number of letters, numbers, and symbols, storing them in the map
	//
	for _, word := range words {
		for _, char := range word {
			if(unicode.IsLetter(char)){
				count["letter"]++
			} else if (unicode.IsNumber(char)) {
				count["number"]++
			} else {
				count["symbol"]++
			}
		}
	}

	// dump the map to the console using the spew package
	//
	spew.Dump(count)
}
