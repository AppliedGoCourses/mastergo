package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

// Remember to always pass wg as a pointer! Copying the wait group
// defeats its purpose.
func worker(n int, wg *sync.WaitGroup) {

	// Decrease the wg count whenever the goroutine exits
	defer wg.Done()

	for i := 0; i < rand.Intn(10)+10; i++ {
		fmt.Print(n)
		time.Sleep(100 * time.Millisecond)
	}
}

func main() {

	// Create a wait group
	var wg sync.WaitGroup

	for i := 0; i < 10; i++ {
		// For each goroutine spawned, add 1 to the wait group.
		// Always call Add() in the calling goroutine, never in the called one. Within the called goroutine, Add() may run too late, and wg.Wait() might then run into a negative goroutine count.
		wg.Add(1)
		go worker(i, &wg)
	}

	// Wait for all goroutines to finish

	wg.Wait()
	fmt.Println("\nDone.")
}
