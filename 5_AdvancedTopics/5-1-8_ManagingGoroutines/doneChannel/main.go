package main

import (
	"fmt"
	"time"
)

func main() {

	fmt.Println("*** Wait until a goroutine ends ***")

	// This is BTW a good way to define local helper functions
	// that are used by the enclosing function only.
	doneFunc := func(data <-chan int, done chan<- struct{}) {
		// Produce some data
		for i := 0; i < 10; i++ {
			fmt.Println(<-data * 2)
			time.Sleep(time.Microsecond)
		}
		// Let the main goroutine know we're done
		close(done) // also possible: done <- struct{}{}
	}

	dataChan := make(chan int, 10)
	doneChan := make(chan struct{})

	go doneFunc(dataChan, doneChan)

	for i := 0; i < 10; i++ {
		dataChan <- i * 10
	}

	// Wait for the doneChan channel to close
	<-doneChan
	fmt.Println("Done.")

	fmt.Println("\n*** Tell a goroutine to end ***")

	worker := func(stop <-chan struct{}) {
		i := 0
		for {
			select {
			case <-stop:
				return
			default:
				fmt.Print(i, " ")
				i++
				time.Sleep(10 * time.Millisecond)
			}
		}
	}

	stopChan := make(chan struct{})

	go worker(stopChan)

	time.Sleep(time.Second)

	close(stopChan)
	fmt.Println("Done.")
}
