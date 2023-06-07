// concurrency/channel-ranging/begin/main.go
package main

import (
	"fmt"
	"math/rand"
	"time"
)

// fill a channel with ints up to max
func fill(ch chan<- int, max int) {
	// ensure non-deterministic random numbers
	s := rand.NewSource(time.Now().UnixNano())
	r := rand.New(s)

	// loop and fill the channel up to its capacity qwith random numbersyy
	for i := 0; i < cap(ch); i++ {
		ch <- r.Intn(max)
	}

	// close the channel and signal that no more values will be sent
	close(ch)
}

func main() {
	// create a channel with a buffer size
	channel := make(chan int, 5)

	// invoke the fill function as a goroutine
	fill(channel, 100)

	// range over the channel and print out the values
	for n := range channel {
		fmt.Println(n)
	}
}
