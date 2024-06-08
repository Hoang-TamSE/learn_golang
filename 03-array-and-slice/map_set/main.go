package main

import "fmt"

func main() {
	fmt.Println("Map set")
	map_set()
	fmt.Println("Map set struct")
	map_set_struct()
}

func map_set() {
	intSet := map[int]bool{}
	vals := []int{5, 10, 2, 5, 8, 7, 3, 9, 1, 2, 10}

	for _, v := range vals {
		intSet[v] = true
	}

	fmt.Println(len(vals), len(intSet))
	fmt.Println(intSet[5])
	fmt.Println(intSet[500])

	if intSet[100] {
		fmt.Println("100 is in the set")
	}
}

func map_set_struct() {

	intSet := map[int]struct{}{}
	vals := []int{5, 10, 2, 5, 8, 7, 3, 9, 1, 2, 10}

	for _, v := range vals {
		intSet[v] = struct{}{}
	}

	if _, ok := intSet[5]; ok {
		fmt.Println("5 is in the set")
	}
}
