// concurrency/channel-basics/begin/main.go
package main

import (
	"fmt"
)

// sum calculates and prints the sum of numbers
func sum(nums []int, ch chan<- int) {
	sum := 0
	for _, v := range nums {
		sum += v
	}
	ch <- sum
}

func main() {
	nums := []int{1, 2, 3, 4, 5}

	channel := make(chan int)

	// invoke the sum function as a goroutine
	go sum(nums, channel)

	result := <-channel

	// force main thread to sleep
	fmt.Println(result)

	channelbuffered := make(chan string, 2)
	channelbuffered <- "Jon"
	channelbuffered <- "Paul"
	fmt.Println(<-channelbuffered)
	fmt.Println(<-channelbuffered)
}
