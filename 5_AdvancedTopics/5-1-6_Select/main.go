package main

import (
	"fmt"
)

// Problem: all channel operations are serialized. If one channel operation blocks,
// the whole code blocks.
func forLoop(c1, c2, c3, c4 chan string) {
	for {
		fmt.Println(<-c1)
		s := <-c3
		c2 <- s
		if s := <-c4; s == "END" {
			return
		}
	}
}

// Solution: select. Each case can run as soon as its channel operation succeeds.
func selectLoop(c1, c2, c3, c4 chan string) {
	for {
		select {
		case s := <-c1:
			fmt.Println(s)
		case s := <-c3:
			c2 <- s
		case s := <-c4:
			if s == "END" {
				return
			}
		}
	}
}

func main() {
	c1 := make(chan string)
	c2 := make(chan string)
	c3 := make(chan string)
	c4 := make(chan string)
	done := make(chan bool)

	go selectLoop(c1, c2, c3, c4)

	go func(c chan string, done chan bool) {
		fmt.Println("From c2:", <-c)
		done <- true
	}(c2, done)

	c4 <- "BEGIN"
	c1 <- "Message from c1"
	c3 <- "From c3"
	c4 <- "END"
	<-done
}
