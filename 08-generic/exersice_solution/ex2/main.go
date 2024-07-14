package main

import (
	"fmt"
	"strconv"
)

type Printable interface {
	~int | ~float64
	fmt.Stringer
}

type PrintInt int

func (pi PrintInt) String() string {
	return strconv.Itoa(int(pi))
}

type PrintFloat float64

func (pf PrintFloat) String() string {
	return fmt.Sprintf("%f", pf)
}

func PrintIt[T Printable](v T) {
	fmt.Println(v)
}

func main() {
	var i PrintInt = 20
	PrintIt(i)

	var f PrintFloat = 10.23
	PrintIt(f)
}
