// flow-control/loop-range/begin/main.go
package main

import "fmt"

func main() {
	// initialize a slice of ints
	test := []int{1, 2, 3, 4, 5}

	// use for-range to iterate over the slice and print each value
	for i, v := range test {
		fmt.Println(i, v)
	}

	// declare a map of strings to ints
	//
	testMap := map[string]int{"a": 1, "b": 2}

	// use for-range to iterate over the map and print each key/value pair
	for i, v := range testMap {
		fmt.Println(i, v)
	}
}
