package main

import (
	"fmt"
	"maps"
)

func main() {
	deleteMap()
	emptyMap()
	comparingMap()
}

func deleteMap() {
	m := map[string]int{
		"hello": 5,
		"world": 10,
	}
	fmt.Println(m)
	delete(m, "hello")
	fmt.Println(m)
}

func emptyMap() {
	m := map[string]int{
		"hello": 5,
		"world": 10,
	}
	fmt.Println(m, len(m))
	clear(m)
	fmt.Println(m, len(m))

}

func comparingMap() {
	m := map[string]int{
		"hello": 5,
		"world": 10,
	}

	n := map[string]int{
		"world": 10,
		"hello": 5,
	}

	fmt.Println(maps.Equal(m, n)) // true
}
