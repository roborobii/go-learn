package main

import (
	"fmt"
)

// Worker Pulls (common concurrency pattern)
// 	where a queue of work to be done and multiple concurrent workers
// 	pulling off items from queue

func main() {
	jobs := make(chan int, 100) // buffer channel of 100
	results := make(chan int, 100)

	// makes use of
	go worker(jobs, results)
	go worker(jobs, results)
	go worker(jobs, results)
	go worker(jobs, results)

	for i := 0; i < 50; i++ {
		jobs <- i
	}
	close(jobs)

	for j := 0; j < 50; j++ {
		fmt.Println(<-results)
	}

}

// only receive on jobs channel and only send on results channel
func worker(jobs <-chan int, results chan<- int) {
	for n := range jobs {
		results <- fib(n)
	}
}

func fib(n int) int {
	// if n == 0 || n == 1 {
	// 	return n
	// }
	// first, second := 0, 1
	// for i := 2; i <= n; i++ {
	// 	temp := first + second
	// 	first = second
	// 	second = temp
	// }
	// return second

	// inefficient fib to demonstrate
	if n <= 1 {
		return n
	}
	return fib(n-1) + fib(n-2)
}
