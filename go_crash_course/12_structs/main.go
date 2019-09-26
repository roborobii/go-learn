package main

import (
	"fmt"
	"strconv"
)

// Person struct represents a person
type Person struct {
	// firstName string
	// city string
	// age int

	firstName, city string
	age             int
}

// Method (value receiver - doesn't change anything)
func (p Person) greet() string {
	return "Hello, my name is " + p.firstName + " " + "and I am " + strconv.Itoa(p.age)
}

// Method (pointer receiver - sets/changes something)
func (p *Person) hasBirthday() {
	p.age++
}
func (p *Person) nameChange(name string) {
	p.firstName = name
}

func main() {
	// Initialize Person using struct
	person1 := Person{firstName: "Robin", age: 22, city: "San Jose"}
	person2 := Person{"Robin", "San Jose", 22} // alternative
	fmt.Println(person1)
	fmt.Println(person2)
	fmt.Println(person2.firstName)

	fmt.Println(person1.greet())
	person1.hasBirthday()
	fmt.Println(person1.greet())
	person1.nameChange("Rob")
	fmt.Println(person1.greet())
}
