package main

import (
	"fmt"
	"runtime"
	"time"
)

// *** Caveat 1: Spawning goroutines in a loop

func spawnInALoop() {
	for i := 0; i < 10; i++ {
		// Our goroutine is a closure, and closures can
		// access variables in the outer scope, so we can
		// grab i here to give each goroutine an ID, right?
		// (Hint: wrong.)
		go func() {
			fmt.Println("Goroutine", i)
		}() // <- Don't forget to call () the closure
	}
	time.Sleep(100 * time.Millisecond)
}

func spawnInALoopFixed() {
	for i := 0; i < 10; i++ {
		// Always pass any start value properly as a
		// function parameter. This way, nothing can go wrong.
		go func(n int) {
			fmt.Println("Goroutine", n)
		}(i) // We pass i here as an argument
	}
	time.Sleep(100 * time.Millisecond)
}

// *** Caveat 2: Goroutine leak

type Worker struct {
	Ch chan string
}

func (w *Worker) work() {
	for {
		w.Ch <- "worker here!"
	}
}

func NewWorker() *Worker {
	w := &Worker{Ch: make(chan string, 100)}
	go w.work()
	return w
}

func goroutineleak() {
	w := NewWorker()
	fmt.Println(<-w.Ch)
	// After this point, w goes out of scope
	// but the goroutine continues to exist
	// (and so does w). Goroutine leak!
}

func main() {
	fmt.Println("*** Spawn in a loop ***")
	spawnInALoop()

	fmt.Println("\n*** Fixed ***")
	spawnInALoopFixed()

	fmt.Println("\n*** Goroutine leak ***")

	fmt.Println(runtime.NumGoroutine())
	goroutineleak()
	fmt.Println(runtime.NumGoroutine())

	fmt.Println("\n***  ***")
}
