// The solution to task #2: How to use channel axiom 4 for signaling that work has finished

package main

import (
	"fmt"
)

func worker(done chan<- struct{}) {
	// Do something for some time
	for i := 0; i < 1000; i++ {
		fmt.Print(".")
	}
	fmt.Println()
	// Signal end of work by closing the channel
	close(done)
}

func main() {
	// done is an unbuffered channel of empty structs. Any other data
	// type would do as well, but the empty struct makes clear that we
	//  are not interested to send any data through the channel.
	done := make(chan struct{})

	// start the goroutine.
	go worker(done)

	// Now wait for the goroutine to finish by simply attempting to
	// read from the done channel. According to axiom 4, the read
	// attempt blocks until the channel gets closed and starts delivering
	// zero values.
	fmt.Println("Waiting for the goroutine")
	<-done
	fmt.Println("Done")

	// Also test what's happening when commenting out the <-done statement.

	// Note that "<-done" discards the value read from the channel.
	// The compiler does not complain if we do not assign the result
	// to a variable. No blank identifier needed here.

}
