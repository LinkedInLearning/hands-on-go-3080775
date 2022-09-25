package main

import "fmt"

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

	// create a channel
	ch := make(chan int)

	// invoke sum as a goroutine and pass the channel
	go sum(nums, ch)

	// receive the result from the channel
	result := <-ch // blocking operation

	fmt.Println("Result:", result)

	// create a buffered channel
	ch2 := make(chan string, 2)

	// add two values to the channel
	ch2 <- "James"
	ch2 <- "Toni"

	// receive the result from the channel
	fmt.Println(<-ch2)
	fmt.Println(<-ch2)
}
