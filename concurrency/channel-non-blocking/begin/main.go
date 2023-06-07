// concurrency/channel-non-blocking/begin/main.go
package main

import "fmt"

func main() {
	// declare a signal channel to read from
	notreadchannel := make(chan struct{})

	select {
	case <-notreadchannel:
		fmt.Println("channel1")
	default:
		fmt.Println("Nothing was received")

	}

	// use a default case in select to prevent blocking
}
