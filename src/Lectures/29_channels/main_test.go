package main

import "testing"

func BenchmarkChanZeroLength(b *testing.B) {
	for i := 0; i < b.N; i++ {
		ints := make(chan int, 0)
		go produceSteadily(ints)
		consumeInBatches(ints)
	}
}

func BenchmarkChanLengthTen(b *testing.B) {
	for i := 0; i < b.N; i++ {
		ints := make(chan int, 10)
		go produceSteadily(ints)
		consumeInBatches(ints)
	}
}

// Results:
//
// BenchmarkChanZeroLength-8              1        2104104307 ns/op
// BenchmarkChanLengthTen-8               1        1145659979 ns/op
