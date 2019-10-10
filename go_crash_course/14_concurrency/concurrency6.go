package main

import (
	"fmt"
	"time"
)

// Select: receiving/executing from ready channels
func main() {
	c1 := make(chan string)
	c2 := make(chan string)

	go func() {
		for {
			c1 <- "Every 500ms"
			time.Sleep(time.Millisecond * 500)
		}
	}()

	go func() {
		for {
			c2 <- "Every 2s"
			time.Sleep(time.Second * 2)
		}
	}()

	// instead of this: (which always waits/blocks for the slowest)
	// for {
	// 	fmt.Println(<-c1)
	// 	fmt.Println(<-c2)
	// }

	// do this: (use select statement to receive from ready channels)
	for {
		select {
		case msg1 := <-c1: // whenever c1 is ready, get it
			fmt.Println(msg1)
		case msg2 := <-c2: // whenever c2 is ready, get it
			fmt.Println(msg2)
		}
	}
}
