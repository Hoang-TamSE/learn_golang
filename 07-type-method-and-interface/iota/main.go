package main

import "fmt"

const (
	Field1 = 0
	Field2 = 1 + iota
	Field3 = 20
	Field4
	Field5 = iota
)

func main() {
	fmt.Println(Field1, Field2, Field3, Field4, Field5)
}
