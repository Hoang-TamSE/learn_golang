package main

import (
	"fmt"
	"math/rand"
)

func main() {
	result := make([]int, 0, 100)
	for i := 0; i < 100; i++ {
		result = append(result, rand.Intn(100))
	}
	for _, v := range result {
		switch {
		case v%6 == 0:
			fmt.Println("Six!")
		case v%2 == 0:
			fmt.Println("Two!")
		case v%3 == 0:
			fmt.Println("Three!")
		default:
			fmt.Println("Nerver Mind")
		}

	}
}
