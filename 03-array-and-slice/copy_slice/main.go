package main

import "fmt"

func main() {
	x := []int{1, 2, 3, 4}
	y := make([]int, 4)

	num := copy(y, x)

	fmt.Println(y, num) //[1 2 3 4] 4

	// x := []int{1, 2, 3, 4}
	// y := make([]int, 2)
	// copy(y, x[2:]) [3, 4]

	// x := []int{1, 2, 3, 4}
	// num := copy(x[:3], x[1:])
	// fmt.Println(x, num) [2 3 4 4] 3.

}
