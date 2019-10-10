package main

import (
	"fmt"
	"time"
)

/*
Paralellism: Running 2 or more things at the EXACT same time
Concurrency: Breaking up a program into independently executable tasks
	that could run at the same time while still getting right result
*/

func main() {
	// count("sheep") // will loop with sheep forever
	// count("dog")

	go count("sheep")
	// using go (creates a goroutine); will not wait for execution to finish and move on to next line
	// "go and run this execution in the background and continue to next line immediately"
	// count("dog")
	// the main execution is a goroutine, calling another go will make it 2 goroutines

	go count("awesome")
	// when the main goroutine exits, the program exits no matter what other go routine does
	// time.Sleep(time.Second * 2) // sleep for two seconds
	fmt.Scanln() // waits for an input before exiting out of main goroutine
}

func count(input string) {
	for i := 1; true; i++ {
		fmt.Println(i, input)
		time.Sleep(time.Millisecond * 500)
	}
}
