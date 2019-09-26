package main

import "fmt"

func main() {
	array := [6]int{33, 32, 39, 95, 19, 2}
	// loop through array (or slice) using range (similar to enumerate in python)
	for index, element := range array {
		// fmt.Println(index, element)
		fmt.Printf("index=%d: element=%d\n", index, element)
	}
	// when not using index*
	for _, element := range array {
		// fmt.Println(index, element)
		fmt.Printf("element=%d\n", element)
	}

	emails := map[string]string{"bob": "bob@gmail.com", "rob": "rob@gmail.com"}
	// loop through map using range
	for key, value := range emails {
		fmt.Printf("%s: %s\n", key, value)
	}

	for key := range emails {
		fmt.Printf("%s\n", key)
	}
}
