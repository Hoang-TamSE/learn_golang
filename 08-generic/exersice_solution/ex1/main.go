package main

import "fmt"

type DoubleType interface {
	~int | ~int8 | ~int16 | ~int32 | ~int64 |
		~uint | ~uint8 | ~uint16 | ~uint32 | ~uint64 | ~uintptr |
		~float32 | ~float64
}

func Double[T DoubleType](val T) T {
	return val * 2
}

func main() {
	t1 := 20.222
	fmt.Println(Double(t1))

	fmt.Println(Double(2))
}
