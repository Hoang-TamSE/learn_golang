package main

import "fmt"

func main() {
	var x = [12]int{1, 2, 3, 10: 1, 3}
	fmt.Println(x)

	var y = [...]int{1, 2, 3}
	var z = [3]int{1, 2, 3}
	fmt.Println(y == z)

	var j [2][3]int
	j[0][1] = 1
	print(j[0][1])
}
