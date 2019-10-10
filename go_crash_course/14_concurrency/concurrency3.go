package main

import (
	"fmt"
	"time"
)

// Channels
func main() {
	c := make(chan string) // make a channel
	go count("sheep", c)   // sending in the function

	// for {
	// 	msg, open := <-c // receiving message
	// 	if !open {
	// 		break
	// 	}
	// 	fmt.Println(msg)
	// }
	for msg := range c { // similar function as above
		// but iterates the range of channel, until it is closed
		fmt.Println(msg)
	}

	// sending and receiving are blocking operations
	//	when trying to receive, have to wait for a value there to be received
	// 	when trying to send, have to wait till receiver is ready to receive
	// this blocking nature of channels allow synchronization of goroutines

	// as a sender, we can close the channel when we're finished sending
	// receivers should not close channels bc we don't know if sender is done
}

func count(thing string, c chan string) {
	for i := 1; i <= 5; i++ {
		c <- thing // sending a message into the channel
		time.Sleep(time.Millisecond * 500)
	}
	close(c)
}
