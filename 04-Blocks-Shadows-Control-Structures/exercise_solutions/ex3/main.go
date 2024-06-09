package main

import "fmt"

func main() {
	total := 0
	for i := 0; i < 10; i++ {
		//should be = not :=
		total := total + i
		fmt.Println(total)
	}
	fmt.Println("total:", total)
	/*
		0
		1
		2
		3
		4
		5
		6
		7
		8
		9
		total: 0
		value of total is 0 because shadown variable
	*/
}
