package main

import (
	"fmt"
	"sync"
	"sync/atomic"
)

func race(n int) {
	// This counter is going to be used by multiple goroutines - data race ahead!
	var counter int

	wg := &sync.WaitGroup{}

	// This is our goroutine code. It initialized the counter if it is still nil,
	// and then increments it 10 times.
	count := func(wg *sync.WaitGroup) {
		for i := 0; i < 1000; i++ {
			counter++
		}
		wg.Done()
	}

	// We spawn n goroutines now.
	wg.Add(n)
	for i := 0; i < n; i++ {
		go count(wg)
	}
	wg.Wait()

	// The counter's final value should now be n * 10.
	fmt.Println("Final value:", counter)
}

func mutex(n int) {
	var counter int64

	// A mutex guards the execution of statements.
	var m sync.Mutex

	wg := &sync.WaitGroup{}

	count := func(wg *sync.WaitGroup) {

		for i := 0; i < 100; i++ {
			// A goroutine that locks a mutex has exclusive access
			// to shared data.
			// A goroutine that tries to lock an already locked mutex
			// must wait until the mutex is unlocked.
			m.Lock()

			counter++

			// When finished with accessing the shared data, unlock the mutex.
			// Failing to unlock properly can cause deadlocks.
			m.Unlock()
		}
		wg.Done()
	}

	wg.Add(n)
	for i := 0; i < n; i++ {
		go count(wg)
	}
	wg.Wait()
	fmt.Println("Final value:", counter)
}

func atomicInc(n int) {
	var counter int64

	wg := &sync.WaitGroup{}

	count := func(wg *sync.WaitGroup) {

		for i := 0; i < 100; i++ {

			// atomic contains some primitives for atomically handling
			// integer operations like add or swap.
			// Main usage is the construction of higher-level synchronization
			// mechanisms.
			// In most cases, channels or mutexes are preferred over atomic updates.
			atomic.AddInt64(&counter, 1)
		}
		wg.Done()
	}

	wg.Add(n)
	for i := 0; i < n; i++ {
		go count(wg)
	}
	wg.Wait()
	fmt.Println("Final value:", counter)
}

func once(n int) {
	var data map[string]int
	var m sync.Mutex
	var wg sync.WaitGroup

	var once sync.Once

	setup := func() {
		fmt.Println("In setup()")
		data = make(map[string]int, n)
	}

	worker := func(s string, n int, wg *sync.WaitGroup) {
		// No matter how many goroutines call this,
		// it only gets executed once.
		fmt.Println("Calling once.Do(setup)")
		once.Do(setup)

		m.Lock()
		data[s] = n
		m.Unlock()

		wg.Done()
	}

	wg.Add(n)
	for i := 0; i < n; i++ {
		go worker(fmt.Sprintf("go%d", i), i*i, &wg)
	}

	wg.Wait()
	fmt.Printf("Map 'data': %v\n", data)
}

func main() {

	fmt.Println("*** Data race ***")
	for i := 0; i < 10; i++ {
		race(10)
	}

	fmt.Println("\n*** Mutex ***")
	for i := 0; i < 10; i++ {
		mutex(10)
	}

	fmt.Println("\n*** Atomic ***")
	for i := 0; i < 10; i++ {
		atomicInc(10)
	}

	fmt.Println("\n*** Once ***")
	once(10)
}
