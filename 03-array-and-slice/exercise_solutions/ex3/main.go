package main

import "fmt"

type Employee struct {
	firstName string
	lastName  string
	id        int
}

func main() {
	first := Employee{
		"tam",
		"le",
		1,
	}

	second := Employee{
		firstName: "My",
		lastName:  "Truong",
		id:        2,
	}

	var third Employee
	third.firstName = "Trinh"
	third.lastName = "Khau"
	third.id = 3
	fmt.Println(first)
	fmt.Println(second)
	fmt.Println(third)
}
