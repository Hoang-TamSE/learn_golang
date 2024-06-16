package main

import "fmt"

type Person struct {
	firstName string
	lastName  string
	age       int
}

func main() {
	p := MakePerson("Fred", "Williamson", 25)
	fmt.Println(p)
	p2 := MakePersonPointer("Wilma", "Fredson", 32)
	fmt.Println(p2)
}

func MakePerson(firstName string, lastName string, age int) Person {
	return Person{
		firstName: firstName,
		lastName:  lastName,
		age:       age,
	}
}

func MakePersonPointer(firstName string, lastName string, age int) *Person {
	return &Person{
		firstName: firstName,
		lastName:  lastName,
		age:       age,
	}
}
