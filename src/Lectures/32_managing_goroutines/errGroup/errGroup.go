package main

import (
	"context"
	"errors"
	"fmt"
	"log"
	"time"

	"golang.org/x/sync/errgroup"
)

func likeWaitGroupButWithErrors() {

	// Create a wait group
	var g errgroup.Group
	chans := [10]chan int{} // Look Ma, an array!

	for i := 0; i < 10; i++ {

		c := make(chan int, 10)
		chans[i] = c // save for later
		n := i       // capture current value of i for the closure

		// Spawn goroutines the errgroup way
		g.Go(func() error {
			for m := range chans[n] { // use the captured loop index n here
				fmt.Print(n, "-", m, " ")
			}
			return nil
		})
	}

	// Feed the beast^H^H^H^H^H goroutines

	go func() {
		for i := 0; i < 10; i++ {
			for j := 0; j < 10; j++ {
				chans[i] <- j
			}
			close(chans[i])
		}
	}()

	// Wait for all goroutines to finish, check for errors

	err := g.Wait()
	if err != nil {
		log.Fatalln(err)
	}
	fmt.Println("Done.")
}

func errgroupWithContext() {
	c, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	g, ctx := errgroup.WithContext(c)

	for i := 0; i < 10; i++ {
		n := i // capture current loop value for the closure
		g.Go(func() error {
			for {
				select {
				case <-ctx.Done():
					fmt.Print("\nGoroutine ", n, " timed out")
					return errors.New("timeout ocurred")
				default:
					fmt.Print(n, " ")
					time.Sleep(100 * time.Millisecond)
				}
			}
		})
	}

	err := g.Wait()
	fmt.Println()
	fmt.Println("Error returned:", err)
	fmt.Println("Done.")
}

func main() {

	fmt.Println("*** errgroup.group instead of sync.WaitGroup ***")

	likeWaitGroupButWithErrors()

	fmt.Println("\n*** errgroup with context with timeout ***")

	errgroupWithContext()
}
