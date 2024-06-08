package main

import "fmt"

func main() {
	// make function create slice with specify the type, len, and cap

	x := make([]int, 5) // create slice int with len and cap equal 5

	fmt.Println(x)

	y := make([]int, 5)
	y = append(y, 10)

	fmt.Println(y) //[0 0 0 0 0 10]

	z := make([]int, 2, 2)
	z = append(z, 3)
	fmt.Print(z)

	/*
		• If you are using a slice as a buffer (you’ll see this in “io and Friends” on page 319),
		then specify a nonzero length.

		• If you are sure you know the exact size you want, you can specify the length and
		index into the slice to set the values. This is often done when transforming values
		in one slice and storing them in a second. The downside to this approach is that
		if you have the size wrong, you’ll end up with either zero values at the end of the
		slice or a panic from trying to access elements that don’t exist.

		• In other situations, use make with a zero length and a specified capacity. This
		allows you to use append to add items to the slice. If the number of items turns
		out to be smaller, you won’t have an extraneous zero value at the end. If the
		number of items is larger, your code will not panic.

	*/
}
