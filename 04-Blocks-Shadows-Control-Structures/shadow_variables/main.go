package main

import "fmt"

func main() {
	x := 10
	if x > 5 {
		fmt.Println(x) // 10
		var x int = 5
		fmt.Println(x) // 5
	}
	fmt.Println(x) // 10
}
