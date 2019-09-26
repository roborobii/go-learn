package main

import "fmt"

func main() {
	// Map
	emails := make(map[string]string) // make

	// assign key value
	emails["bob"] = "bob@gmail.com"
	emails["rob"] = "rob@gmail.com"

	fmt.Println(emails)
	fmt.Println(emails["bob"])
	fmt.Println(emails["no"])

	delete(emails, "bob") // delete
	fmt.Println(emails)

	// declare and assign key values
	emails2 := map[string]string{"bob": "bob@gmail.com", "rob": "rob@gmail.com"}
	fmt.Println(emails2)
}
