package main

import (
	"fmt"
	"math/rand"
	"time"
)

const (
	dataRate      = 10   // values per interval
	printInterval = 1000 // milliseconds
)

// startDatGenerator starts a goroutine that produces random
// "stock market chart" data at an interval of `dataRate` ms.
// The data is sent to chan data.
func startDataGenerator(data chan<- float64) {
	const max = 60
	const volatility = 0.1
	value := rand.Float64()
	go func() {
		for {
			rnd := 2 * (rand.Float64() - 0.5)
			change := volatility * rnd
			change += (0.5 - value) * 0.1
			value += change

			data <- value * max

			time.Sleep(time.Duration(dataRate * time.Millisecond)) // simulate limited data rate
		}
	}()
}

// Collect data collects the generated data and prints out the
// average of all values received in intervals of `printInterval` milliseconds.
//
// For this it needs to listen to two channels:
// * One that delivers the data, and
// * one that delivers a time tick every second.
//
// Reading on a channel blocks until data is available, so
// we need a select statement here that lets us multiplex
// the channels. Whenever a channel delivers data, the
// corresponding case is invoked.
func collectData(data <-chan float64) {

	// list is a ring buffer, and head holds the current input position.
	list := make([]float64, dataRate, dataRate)
	head := 0

	// The time library has a nice feature: A ticker that
	// sends regular clock ticks down a channel.
	// We use this here to print out the rolling average
	// of the accumulated "stock" data every second.
	tick := time.NewTicker(time.Second)

	for {
		select {

		case val := <-data:
			// Insert the value into the ring buffer.
			list[head] = val
			head = (head + 1) % dataRate // The modulo operator keeps the head position within the borders of the slice

		case <-tick.C:
			// Calculate the average value from the current slice contents
			sum := 0.0
			for _, v := range list {
				sum += v
			}
			avg := sum / float64(dataRate)

			// Print the avg value as an "ASCII" diagram
			for i := 0; i < int(avg); i++ {
				fmt.Print(" ")
			}
			fmt.Println("*")
		}
	}
}

func main() {
	dch := make(chan float64)
	startDataGenerator(dch)
	collectData(dch)
}
