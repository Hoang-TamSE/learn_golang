package main

import "fmt"

// this code compiles, but lookalike characters are a very bad idea.
func main() {
	a := "hello"   // standard lowercase a (Unicode U+0061), different from the line above!
	ａ := "goodbye" // Unicode U+FF41
	fmt.Println(a)
	fmt.Println(ａ)

}
