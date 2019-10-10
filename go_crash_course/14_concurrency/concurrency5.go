package main

import (
	"fmt"
)

// Buffered channel: Blocking solution, can fill up to a capacity before blocking
func main() {
	c := make(chan string, 2)
	c <- "hello"
	c <- "world"

	msg := <-c
	fmt.Println(msg)
	msg = <-c
	fmt.Println(msg)
}
