package main

import "fmt"

func main() {
	x := 10
	fmt.Println(x)
	fmt := "oops"
	fmt.Println(fmt)
	//.\main.go:9:6: fmt.Println undefined (type string has no field or method Println)
}
