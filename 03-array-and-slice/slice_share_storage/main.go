package main

import "fmt"

// Slices with overlapping storage
func main() {
	x := []string{"a", "b", "c", "d"}
	y := x[:2] // [a, b]
	z := x[1:] // [b, c, d]

	x[1] = "y" // x = [a, y, c, d]
	y[0] = "x" // x = [x, y, c, d]
	z[1] = "z" // x [x, y, z, d]

	fmt.Println("x:", x) // x: [x y z d]
	fmt.Println("y:", y) // y: [x y]
	fmt.Println("z:", z) // z: [y z d]

	//Changing x modified both y and z, while changes to y and z modified x.

}
