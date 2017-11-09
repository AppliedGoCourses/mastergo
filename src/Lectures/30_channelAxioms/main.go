package main

import (
	"fmt"
	"os"
	"time"
)

func nilChannel() {
	var c chan int // the zero value is nil

	// This is our sender goroutine.
	go func(c chan int) {
		// Try to send something to a nil channel.
		c <- 42
	}(c)

	// Now try receiving something from the nil channel.
	n := <-c
	fmt.Println(n)
}

func sendToClosedChannel(c chan int) {
	c <- 1
	c <- 2
	c <- 3
	close(c)
	// wait a bit to let the receiver read some zero values
	time.Sleep(100 * time.Microsecond) // adjust this if your computer is faster or slower than mine
	// Now try to send something through the closed channel.
	c <- 4
}

func receiveFromClosedChannel(c chan int) {
	for {
		n := <-c
		fmt.Println("Read:", n)
	}
}

func receiveFromClosedChannelCommaOk(c chan int) {
	for {
		n, ok := <-c
		if !ok {
			fmt.Println("Comma,ok: the channel is closed")
			return
		}
		fmt.Println("Comma,ok: read", n)
	}
}

func main() {
	// Since both tests trigger a fatal error, you have to decide which
	// test to run. Pass "nil" as an argument in order to run the nil channel test,
	// and no argument in order to run the closed channel tests.

	if len(os.Args) > 1 && os.Args[1] == "nil" {
		fmt.Println("A send to a nil channel blocks forever. ")
		fmt.Println("A receive from a nil channel blocks forever. ")
		nilChannel()
	}

	fmt.Println("A send to a closed channel panics.")
	fmt.Println("A receive from a closed channel returns the zero value immediately.")
	c := make(chan int, 10)
	// Sorry for the unwieldy function names...
	go sendToClosedChannel(c)
	go receiveFromClosedChannel(c)
	receiveFromClosedChannelCommaOk(c)
	time.Sleep(time.Second)

}
