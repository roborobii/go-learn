package main

import "fmt"

func main() {
	/*
		Main Data Types:
		string, bool
		int, int8, int16, int32, int64
		uint, unit8, uint16, uint32, uint64, uintptr
		byte - alias for uint8
		rune - alias for int32
		float32, float64,
		complex64, complex128

		Must use a variable that you create, else error
	*/
	var name string = "Robin" // type can be inferred string
	var age int32 = 23
	var canBeReassigned = true // type is inferred bool
	canBeReassigned = false
	const cannotBeReassigned = true

	// Shorthand; NOTE: cannot be used outside of function
	name2 := "Jonah" // inferred as string
	size := 1.3 // inferred as float64

	name3, email := "Robin", "robin@gmail.com"

	fmt.Println(name, age, canBeReassigned, cannotBeReassigned)
	fmt.Println(name3, email) // all variables must be used
	fmt.Printf("%T\n", name2) // to get type, use %T
	fmt.Printf("%T\n", size)
}

