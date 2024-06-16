package main

import (
	"fmt"
	"time"
)

type Counter struct {
	total      int
	lastUpdate time.Time
}

func (c *Counter) Increment() {
	c.total++
	c.lastUpdate = time.Now()
}

func (c Counter) String() string {
	return fmt.Sprintf("total: %d, last updated: %v", c.total, c.lastUpdate)
}

func main() {
	main1()
	main2()
}

func main1() {
	var c Counter
	fmt.Println(c.String())
	c.Increment()
	fmt.Println(c.String())
}

func main2() {
	c := &Counter{}
	fmt.Println(c.String())
	c.Increment()
	fmt.Println(c.String())
}
