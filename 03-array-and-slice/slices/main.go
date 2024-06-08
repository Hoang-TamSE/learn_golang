package main

import "fmt"

func main() {

	var x = []int{1, 2, 3}
	fmt.Println(x)

	var y = []int{1, 5: 4, 6, 10: 100, 15}
	fmt.Println(y)

	var dim2 = [][]int{{1, 2, 4}, {3, 4, 4}} //  simulate multidimensional
	fmt.Println(dim2[1][2])
}
