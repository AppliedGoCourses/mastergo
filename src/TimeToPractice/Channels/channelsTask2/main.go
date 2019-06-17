package main

import (
	"fmt"
)

func worker( /* TODO: Maybe receive a channel here? */ ) {
	// Do something for some time
	for i := 0; i < 1000; i++ {
		fmt.Print(".")
	}
	fmt.Println()
	// TODO: How to tell main() that work is done?
}

func main() {

	// start the goroutine.
	go worker( /* TODO: Maybe pass a channel here? */ )

	fmt.Println("Waiting for the goroutine")
	// TODO: Add code to wait for the goroutine.
	// time.Sleep() doesn't count.
	fmt.Println("Done")
}
