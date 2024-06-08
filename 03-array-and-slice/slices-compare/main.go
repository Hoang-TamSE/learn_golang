package main

import (
	"fmt"
	"slices"
)

func main() {

	var x []int           // when declare it will be nil
	fmt.Println(x == nil) // true

	j := []int{1, 2, 3, 4, 5}
	y := []int{1, 2, 3, 4, 5}
	z := []int{1, 2, 3, 4, 5, 6}
	// s := []string{"a", "b", "c"}
	fmt.Println(slices.Equal(j, y)) // prints true
	fmt.Println(slices.Equal(j, z)) // prints false
	// fmt.Println(slices.Equal(j, s)) // does not compile
}
