package main

import (
	"fmt"
	"time"
)

// Remember that if more than one channel operation is ready to
// run, the runtime selects one of them on a random basis.
// You sould therefore see a quite random stream of 0's and 1's
// in the terminal.

func send(c chan int) {

	// Let's stop after a second.
	timeout := time.NewTicker(time.Second)

	for {
		select {
		case c <- 0:
		case c <- 1:
		case <-timeout.C:
			return
		}
	}
}

func main() {
	c := make(chan int) // also try with a buffered channel

	go send(c)

	for {
		fmt.Print(<-c, " ")
	}
}
