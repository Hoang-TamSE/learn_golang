package main

import (
	"fmt"
	"math"
)

func main() {
	var b byte = math.MaxUint8
	var small int32 = math.MaxInt32
	var bigI uint64 = math.MaxUint64

	b = b + 1
	small += 1
	bigI += 1

	fmt.Println(b, small, bigI)
}
