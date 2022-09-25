// concurrency/channel-select/begin/main.go
package main

import (
	"fmt"
	"time"
)

func main() {
	// declare an empty struct channel for signaling
	sigChan := make(chan struct{})

	// declare a timer channel
	timeChan := time.After(2 * time.Second)

	// launch a goroutine to signal after 1 second
	go func() {
		time.Sleep(1 * time.Second)
		sigChan <- struct{}{}
	}()

	// wait for a signal on either channel
	select {
	case <-sigChan:
		fmt.Println("received from sigChan")
	case <-timeChan:
		fmt.Println("received from timeChan")
	}
}
