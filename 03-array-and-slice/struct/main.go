package main

import "fmt"

type person struct {
	name string
	age  int
	pet  string
}

type firstPerson struct {
	name string
	age  int
}

type secondPerson struct {
	name string
	age  int
}

func main() {
	tam := person{
		age:  2,
		name: "my",
		pet:  "aa",
	}

	fmt.Println(tam)
	anonymousStruct()

	fstPerson := firstPerson{}
	sndPerson := secondPerson{}

	f := firstPerson{
		name: "Bob",
		age:  50,
	}
	var g struct {
		name string
		age  int
	}
	g.name = "Bob"
	g.age = 50

	fmt.Println(f == g)

	fmt.Println(fstPerson == firstPerson(sndPerson))
}

func anonymousStruct() {
	var person1 struct {
		name string
		age  int
		pet  string
	}

	person1.name = "bob"
	person1.age = 50
	person1.pet = "dog"
	fmt.Println(person1)

	pet := struct {
		name string
		kind string
	}{
		name: "Fido",
		kind: "dog",
	}
	fmt.Println(pet)
}
