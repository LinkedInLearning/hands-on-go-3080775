// concurrency/goroutines/begin/main.go
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
	sum(nums)

	// invoke the sum function as a goroutine
	numsb := []int{1, 2, 4}
	go sum(numsb)
	time.Sleep(20000 * time.Millisecond)
}
