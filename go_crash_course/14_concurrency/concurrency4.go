package main

import (
	"fmt"
)

// Sending and Receiving through Channels: Blocking operations demonstration
func main() {
	c := make(chan string)
	c <- "hello" // deadlocked, send will block till something is ready to receive
	// but it is never ready to received because we are blocked at sent 'c <- "hello"' in the same routine
	msg := <-c
	fmt.Println(msg)
}
