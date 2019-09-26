package main

import "fmt"

func main() {
	// pointer allows you to point to memory address/location of a value

	a := 5
	b := &a // b now represents memory address of a

	fmt.Println(a, b)

	fmt.Printf("%T\n", a) // int
	fmt.Printf("%T\n", b) // *int - represents a pointer

	// Use * to read val from address
	fmt.Println(*b)  // gives the value at the memory address b
	fmt.Println(*&a) // gives the value at the memory address of a

	*b = 10 // changing the value, wherever the memory address b points to
	// so a would be changed
	fmt.Println(a)

	// pointers: used when have to pass data stored at an address
	// 	passing by the address can increase performance

	// everything in Golang is pass by value so you can use pointers to get around that
}
