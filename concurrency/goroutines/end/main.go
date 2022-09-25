package main

import (
	"fmt"
	"time"
)

// sum calculates and prints the sum of numbers
func sum(nums []int) {
	sum := 0
	for _, v := range nums {
		sum += v
	}
	fmt.Println("Result:", sum)
}

func main() {
	nums := []int{1, 2, 3, 4, 5}
	// sum(nums)

	// invoke the sum function as a goroutine
	go sum(nums)

	// force main thread to sleep
	time.Sleep(100 * time.Millisecond)
}
