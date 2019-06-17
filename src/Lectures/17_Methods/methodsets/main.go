package main

import "fmt"

type Counter int

func (c *Counter) Up() {
	(*c)++
}

func (c *Counter) Reset() {
	*c = 0
}

func (c Counter) Get() int {
	return int(c)
}

func main() {
	var c Counter

	(&c).Up()
	c.Up() // implicitly changed to (&c).Up()

	fmt.Println((&c).Get()) // implicitly changed to c.Get()
	fmt.Println(c.Get())
}
