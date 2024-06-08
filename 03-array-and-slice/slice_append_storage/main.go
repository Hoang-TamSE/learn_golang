package main

import "fmt"

// Slices with overlapping storage
func main() {
	x := []string{"a", "b", "c", "d"}
	y := x[:2]

	fmt.Println(cap(x), cap(y)) // 4 4

	y = append(y, "z")

	fmt.Println("x:", x) //x: [a b z d]
	fmt.Println("y:", y) //y: [a b z]
	/*
		When you make the y slice from x, the length is set to 2, but the capacity is set to 4,
		the same as x. Since the capacity is 4, appending onto the end of y puts the value in
		the third position of x.
	*/
}
