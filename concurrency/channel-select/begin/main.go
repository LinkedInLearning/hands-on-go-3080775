// concurrency/channel-select/begin/main.go
package main

import (
	"fmt"
	"time"
)

func main() {
	// declare an empty struct channel for signaling
	channelone := make(chan struct{})

	// declare a timer channel
	channeltwo := time.After(1 * time.Second)

	// launch a goroutine to signal after 1 second

	go func() {
		time.Sleep(2 * time.Second)
		channelone <- struct{}{}
	}()

	// wait for a signal on either channel
	//
	select {
	case <- channelone : fmt.Println("channelone")
	case <- channeltwo : fmt.Println("channeltwo")
	}
}
