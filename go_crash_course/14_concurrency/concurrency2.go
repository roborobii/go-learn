package main

import (
	"fmt"
	"sync"
	"time"
)

// Wait Group
func main() {
	// wait group instead of fmt.Scanln() which waits for user input
	var wg sync.WaitGroup
	wg.Add(1)
	go func() {
		count("sheep")
		wg.Done() // Done will decrement counter by 1
	}()
	wg.Wait() // Wait will block until counter is 0

}

func count(input string) {
	for i := 1; i < 5; i++ {
		fmt.Println(i, input)
		time.Sleep(time.Millisecond * 500)
	}
}
