package main

import "fmt"

func main() {
	// Arrays - fixed length
	var fruitArr [2]string
	fruitArr[0] = "apple"
	fruitArr[1] = "orange"
	fmt.Println(fruitArr)
	fmt.Println(fruitArr[1])

	// declaring and assigning
	vegieArr := [2]string{"potato", "brocolli"}
	fmt.Println(vegieArr[1])

	// Slices - no specified length
	vegieSlice := []string{"potato", "brocolli", "lettuce"}
	fmt.Println(vegieSlice[1])
	fmt.Println(len(vegieSlice))
}