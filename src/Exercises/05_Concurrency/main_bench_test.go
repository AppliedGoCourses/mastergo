package main

import (
	"log"
	"sync"
	"testing"
)

var once = sync.Once{}
var size = 1000

func setup(b *testing.B) {
	*rows = size
	*cols = size
	once.Do(func() {
		err := generate("benchmark", *rows, *cols)
		if err != nil {
			b.Fatal(err)
		}
	})
}

func BenchmarkMakeTable(b *testing.B) {
	for n := 0; n < b.N; n++ {
		makeTable(size, size)
	}
}
func BenchmarkRead(b *testing.B) {
	setup(b)
	b.ResetTimer()
	for n := 0; n < b.N; n++ {
		_, err := readFromFile("benchmark")
		if err != nil {
			b.Fatal(err)
		}
	}
}

func BenchmarkProcess(b *testing.B) {
	setup(b)
	t, err := readFromFile("benchmark")
	b.ResetTimer()
	if err != nil {
		b.Fatal(err)
	}
	for n := 0; n < b.N; n++ {
		_ = process(t)
	}
}

func BenchmarkWrite(b *testing.B) {
	setup(b)
	t, err := readFromFile("benchmark")
	if err != nil {
		b.Fatal(err)
	}
	res := process(t)
	b.ResetTimer()
	for n := 0; n < b.N; n++ {
		writeToFile("benchmarkstats", res)
	}
}

func BenchmarkAll(b *testing.B) {
	for n := 0; n < b.N; n++ {
		data, err := readFromFile("benchmark")
		if err != nil {
			log.Fatalln(err)
		}

		stats := process(data)
		if err != nil {
			log.Fatalln(err)
		}

		err = writeToFile("benchmarkstats", stats)
		if err != nil {
			log.Fatalln(err)
		}
	}
}
