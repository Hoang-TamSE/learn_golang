package main

import "fmt"

func main() {
	// firsSample()
	// secondExample()
	// thirdExample()
	// fourthExample()
	fifthExample()
}

func firsSample() {
	var x int32 = 10
	var y bool = true
	pointerX := &x
	pointerY := &y
	var pointerZ *string
	fmt.Println("pointerX:", pointerX)
	fmt.Println("pointerY:", pointerY)
	fmt.Println("pointerZ:", pointerZ)
}

func secondExample() {
	x := 10
	pointerToX := &x
	fmt.Println(pointerToX)  // prints a memory address
	fmt.Println(*pointerToX) // prints 10
	z := 5 + *pointerToX
	*pointerToX = 1
	fmt.Println("pointerToX:", x)
	fmt.Println(z) // prints 15
}

func thirdExample() {
	var x *int
	var y *float32
	var z *string
	fmt.Println(x == nil)
	fmt.Println(&x)
	fmt.Println(&y)
	fmt.Println(&z)
	// fmt.Println(*x)
}

func fourthExample() {
	var x = new(int)
	fmt.Println(x == nil) // prints false
	fmt.Println(*x)       // prints 0
}

type person struct {
	FirstName  string
	MiddleName *string
	LastName   string
}

func stringp(s string) *string {
	return &s
}

func fifthExample() {
	y := "Perry123"
	p := person{
		FirstName:  "Pat",
		MiddleName: &y, // This works
		LastName:   "Peterson",
	}
	fmt.Println(*(p.MiddleName))
}
