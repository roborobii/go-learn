package main

import "fmt"

func main() {
	fmt.Println("Hello World!")
	// While (for loop long method)
	i := 1
	for i <= 10 {
		// fmt.Println(i)
		i++
	}

	// For (for loop short)
	for i := 1; i <= 10; i++ {
		// fmt.Println(i)
	}

	// FizzBuzz
	for i := 1; i <= 100; i++ {
		if i%3 == 0 && i%5 == 0 {
			fmt.Println("fizzbuzz")
		} else if i%3 == 0 {
			fmt.Println("fizz")
		} else if i%5 == 0 {
			fmt.Println("buzz")
		} else {
			fmt.Println(i)
		}
	}
}
