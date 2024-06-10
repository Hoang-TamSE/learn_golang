package main

import "fmt"

func main() {
	deferExample()
}

func deferExample() int {
	/*
		You can defer multiple functions in a Go function.
		They run in last-in, first-out
		(LIFO) order; the last defer registered runs first
	*/
	a := 10
	defer func(val int) {
		fmt.Println("first:", val)
	}(a)

	a = 20
	defer func(val int) {
		fmt.Println("second:", a)
	}(a)

	a = 30
	fmt.Println("exting:", a)
	return a
}

func example() {
	defer func() int {
		return 2 // there's no way to read this value
	}()
}
