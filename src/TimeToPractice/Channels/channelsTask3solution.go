package main

import "fmt"

func receiver(n int, c <-chan int) {
	for {
		fmt.Printf("Receiver %d received %d\n", n, <-c)
	}
}

func sender(c chan<- int, done chan<- struct{}) {
	for i := 0; i < 10; i++ {
		c <- i
	}

	// We use the technique from task 2 to signal that the sender has finished.
	close(done)

}

func main() {

	c := make(chan int)

	// The done channel only exists to be closed.
	// We can give it any type, but if you want to save
	done := make(chan struct{})

	go receiver(1, c)
	go receiver(2, c)
	go sender(c, done)

	// Now wait for the sender to finish.
	<-done

	// The receivers are forced to quit when the main program ends.
	//
	// BONUS TASK: Have the sender stop the two receivers by closing
	// channel c before closing the done channel. Make the receivers
	// quit when they detect that c is closed.
}
