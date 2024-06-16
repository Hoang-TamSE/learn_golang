package main

import (
	"fmt"
)

func main() {
	list := []string{"tam", "trinh"}
	fmt.Println(list)
	UpdateSlice(list, "my")
	fmt.Println(list)

	fmt.Println(list)
	GrowSlice(list, "trieu")
	fmt.Println(list)
}

func UpdateSlice(list []string, s string) {
	list[len(list)-1] = s
	fmt.Println("UpdateSlice:", list)
}

func GrowSlice(list []string, s string) {
	list = append(list, s)
	fmt.Println("GrowSlice:", list)
}
