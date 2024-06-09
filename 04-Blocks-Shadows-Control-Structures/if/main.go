package main

import (
	"fmt"
	"math/rand"
)

func main() {
	ifElse()
	ifScope()
	ifBadScope()
}

func ifScope() {
	if n := rand.Intn(10); n == 0 {
		fmt.Println("That's too low")
	} else if n > 5 {
		fmt.Println("That's too big:", n)
	} else {
		fmt.Println("That's a good number:", n)
	}

}

func ifBadScope() {
	if n := rand.Intn(10); n == 0 {
		fmt.Println("That's too low")
	} else if n > 5 {
		fmt.Println("That's too big:", n)
	} else {
		fmt.Println("That's a good number:", n)
	}
	fmt.Println(n) // 04-Blocks-Shadows-Control-Structures\if\main.go:33:14: undefined: n

}

func ifElse() {
	n := rand.Intn(10)

	if n == 0 {
		fmt.Println("That's too low")
	} else if n > 5 {
		fmt.Println("That's too low")
	} else {
		fmt.Println("That's a good number:", n)
	}
}
