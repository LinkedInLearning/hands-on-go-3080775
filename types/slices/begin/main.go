// types/slices/begin/main.go
package main

import (
	"fmt"
)

func main() {
	// declare a slice string the make function
	names := make([]string, 0)

	// append 3 names to the slice
	names = append(names, "John")
	names = append(names, "Jane")
	names = append(names, "Mary")

	// print the slice
	fmt.Println(names)

	// slice the slice using syntax slice[low:high]
	// [Jane Mary]
	fmt.Println(names[1:3])
	// [Jane Mary]
	fmt.Println(names[1:])
	// [John]
	fmt.Println(names[:1])
	// [John Jane Mary]
	fmt.Println(names[:])

	var sliceone []int
	var slicetwo []int = []int{1, 2, 3, 4, 5}

	slicethree := []int{5, 6, 7, 8, 9}
	slicefour := make([]int, 3, 5)
	fmt.Println(sliceone)
	fmt.Println(slicetwo[1:2])
	fmt.Println(slicethree[0:5])
	fmt.Println(slicefour[:])

}
