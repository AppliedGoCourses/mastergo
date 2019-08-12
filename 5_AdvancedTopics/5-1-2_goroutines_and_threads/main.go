package main

// Run 'go get -u -v -t github.com/gosuri/uilive' to install
// realtime terminal output

import (
	"fmt"
	"runtime"
	"time"

	"github.com/gosuri/uilive"
)

func busyWorker() {
	// fmt.Print(".")
	for {
		// always busy! (Until activating the time.Sleep
		// call below, a blocking operation)
		time.Sleep(10 * time.Second)
	}
}

func showGoroutines() {
	term := uilive.New()
	term.Start()
	for {
		n := runtime.NumGoroutine()
		fmt.Fprintln(term, n, "goroutines")
		time.Sleep(1 * time.Second)
	}
}

func main() {

	// Display the current number of goroutines every second.

	go showGoroutines()

	// Spawn large numbers of busy workers
	// Then inspect the running Go binary, to see how many
	// system threads are running.

	for i := 0; i < 100000; i++ {
		go busyWorker()
	}

	time.Sleep(30 * time.Second) // Use Ctrl-C to exit earlier
}
