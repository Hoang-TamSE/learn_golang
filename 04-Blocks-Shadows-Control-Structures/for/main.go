package main

import "fmt"

func main() {
	// forComplete()
	// forMissPartDeclara()
	// forMissPartThird()
	forConditionOnly()
}

func forComplete() {
	for i := 0; i < 10; i++ {
		fmt.Println(i)
	}
}

func forMissPartDeclara() {
	i := 0
	for ; i < 10; i++ {
		fmt.Println(i)
	}
}

func forMissPartThird() {
	for i := 0; i < 10; {
		fmt.Println(i)
		if i%2 == 0 {
			i++
		} else {
			i += 2
		}
	}
}

func forConditionOnly() {
	i := 1
	for i < 100 {
		fmt.Println(i)
		i = i * 2
	}
}

func forWithBreak() {
	for {
		// things to do in the loop
		if !CONDITION {
			break
		}
	}
}

func forWithContinue() {
	for i := 1; i <= 100; i++ {
		if i%3 == 0 && i%5 == 0 {
			fmt.Println("FizzBuzz")
			continue
		}
		if i%3 == 0 {
			fmt.Println("Fizz")
			continue
		}
		if i%5 == 0 {
			fmt.Println("Buzz")
			continue
		}
		fmt.Println(i)
	}
}
