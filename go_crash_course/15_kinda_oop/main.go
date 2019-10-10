package main

import (
	"fmt"
)

/*
This Go program is a simple  illustration of how Go can mimic an OOP
Go lacks inheritance, but supports composition
Go interface provides a kind of polymorphism
Go is not an object-oriented programming (oop) language, 
	but	has features that can mimic an oop language
*/


type Person struct {
	first string
	last  string
	age   int
}

type SecretAgent struct {
	Person            // SecretAgent embeds type Person
	first             string
	securityClearance bool
}

func (p Person) phrase() {
	fmt.Println("Hello, I'm " + p.first + "!")
}

func (s SecretAgent) phrase() {
	fmt.Println("Hello, I'm " + s.first + "!")
}

func main() {
	/*
	In Go:
	1. you don't create classes, you create a type
	2. you don't instantiate, you create a VALUE of a TYPE
	i.e. above you see we created a type: (type Person struct)
		 then create/assign a value, not instantiate
	*/
	p1 := Person{
		first: "Robin",
		last:  "C",
		age:   22,
	}
	p2 := SecretAgent{
		Person{
			"Not Secret Name",
			"Unknown",
			0,
		},
		"Secret",
		true,
	}
	fmt.Println(p1)
	fmt.Println(p2)
	fmt.Println(p2.age)
	fmt.Println(p2.Person.age)
	p1.phrase()
	p2.Person.phrase()
	p2.phrase()

}
