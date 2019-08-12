package main

import (
	"fmt"
	"time"
)

// *** Send and receive

func send(c chan<- string) {
	c <- "Hello"
}

func receive(c <-chan string) {
	var s string
	s = <-c
	fmt.Println(s)
}

func sendAndReceive() {
	c := make(chan string)
	go send(c)
	receive(c)
}

// *** Unbuffered channels synchronize, buffered channels decouple

// The producer needs 10ms for producing and sending one element.
func produceSteadily(c chan<- int) {
	for i := 0; i < 100; i++ {
		fmt.Print("\nsend ", i, " - ")
		c <- i
		time.Sleep(10 * time.Millisecond)
	}
}

// The consumer needs ten elements for his work, but then requires 100ms
// to process these elements.
// The consumer also has a long startup time before reading elements
// from the channel.
// Without a buffered channel, first the producer would be blocked until
// the consumer is ready to receive items; and then the consumer would have
// to slow down to the producer's pace of delivering elements.
func consumeInBatches(c <-chan int) {
	time.Sleep(1000 * time.Millisecond)
	for i := 0; i < 10; i++ {
		// The inner loop consumes from the channel.
		for j := 0; j < 10; j++ {
			fmt.Print("recv ", <-c, " - ")
		}
		fmt.Println()
		// Now wait before reading the next batch.
		time.Sleep(100 * time.Millisecond)
	}
}

func main() {
	fmt.Println("Send and receive")
	sendAndReceive()

	fmt.Println("\n*** Unbuffered channel ***")

	unbufc := make(chan int, 0)
	go produceSteadily(unbufc)
	consumeInBatches(unbufc)

	fmt.Println("\n*** Buffered channel ***")

	bufc := make(chan int, 100)
	go produceSteadily(bufc)
	consumeInBatches(bufc)

}
