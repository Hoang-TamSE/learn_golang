package main

import (
	"fmt"
)

func main() {
	var x = []int{1, 2, 3}
	x = append(x, 4)

	x = append(x, 5, 6, 7)

	y := []int{20, 30, 40}

	x = append(x, y...)

	fmt.Println(x)
}
