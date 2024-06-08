package main

import "fmt"

func main() {
	arrayConversion()
	arrayPointerConversions()
	panicArrayConversions()
}

func arrayConversion() {
	xSlice := []int{1, 2, 3, 4}
	xArray := [4]int(xSlice)

	smallArray := [2]int(xSlice[0:3])
	xSlice[0] = 10

	fmt.Println(xSlice)
	fmt.Println(xArray)
	fmt.Println(smallArray)
}

func arrayPointerConversions() {
	xSlice := []int{1, 2, 3, 4}
	xArrayPointer := (*[4]int)(xSlice)
	xSlice[0] = 10
	xArrayPointer[1] = 20
	fmt.Println(xSlice)
	fmt.Println(xArrayPointer)
}

func panicArrayConversions() {
	xSlice := []int{1, 2, 3, 4}
	panicArray := [5]int(xSlice)
	fmt.Println(panicArray)
	//PS D:\learn_golang\learn_golang\03-array-and-slice\array_conversion> .\array_conversion.exe
}
