package main

import (
	"fmt"
	"math/rand"
)

func main() {
	goodGoto()
	// brokenGoto()
}

func goodGoto() {
	a := rand.Intn(10)
	for a < 100 {
		if a%5 == 0 {
			goto done
		}
		a = a*2 + 1
	}
	fmt.Println("do something when the loop completes normally")
done:
	fmt.Println("do complicated stuff no matter why we left the loop")
	fmt.Println(a)

}

// func brokenGoto() {
// 	a := 10
// 	goto skip
// 	b := 20
// skip:
// 	c := 30
// 	fmt.Println(a, b, c)
// 	if c > a {
// 		goto inner
// 	}
// 	if a < b {
// 	inner:
// 		fmt.Println("a is less than b")
// 	}
// 	/*
// 		04-Blocks-Shadows-Control-Structures\goto\main.go:11:7: goto skip jumps over declaration of b at 04-Blocks-Shadows-Control-Structures\goto\main.go:12:4
// 		04-Blocks-Shadows-Control-Structures\goto\main.go:17:8: goto inner jumps into block starting at 04-Blocks-Shadows-Control-Structures\goto\main.go:19:11
// 	*/
// }
