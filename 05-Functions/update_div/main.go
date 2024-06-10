package main

import (
	"errors"
	"fmt"
	"os"
)

func main() {
	result, remainder, err := divAndRemaider(5, 0)
	if err != nil {
		// fmt.Fprintln(os.Stderr, err)
		fmt.Println(err)
		os.Exit(1)
	}
	fmt.Println(result, remainder)
}

func divAndRemaider(num, denom int) (int, int, error) {
	if denom == 0 {
		return 0, 0, errors.New("cannot divide by zero")
	}
	return num / denom, num % denom, nil
}
