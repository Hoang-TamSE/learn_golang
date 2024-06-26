package main

import "fmt"

type Person struct {
	FirstName string
	LastName  string
	Age       int
}

func (p Person) String() string {
	return fmt.Sprintf("%s %s, age %d", p.FirstName, p.LastName, p.Age)
}

func main() {
	person := Person{
		FirstName: "tam",
		LastName:  "le",
		Age:       22,
	}

	fmt.Println(person.String())
}
