package main

import (
	"log"
	"sync"
	"testing"
)

var once = sync.Once{}
var size = 1000

func setup(b *testing.B) (filename string) {
	filename = fileName("benchmark", size, size)
	once.Do(func() {
		err := generateIfNotExists(filename, size, size)
		if err != nil {
			b.Fatal(err)
		}
	})
	return filename
}

func Benchmark_makeTable(b *testing.B) {
	for n := 0; n < b.N; n++ {
		makeTable(size, size)
	}
}
func Benchmark_read(b *testing.B) {
	fname := setup(b)
	b.ResetTimer()
	for n := 0; n < b.N; n++ {
		_, err := readFromFile(fname, size, size)
		if err != nil {
			b.Fatal(err)
		}
	}
}

func Benchmark_process(b *testing.B) {
	fname := setup(b)
	t, err := readFromFile(fname, size, size)
	b.ResetTimer()
	if err != nil {
		b.Fatal(err)
	}
	for n := 0; n < b.N; n++ {
		_ = process(t)
	}
}

func Benchmark_write(b *testing.B) {
	fname := setup(b)
	t, err := readFromFile(fname, size, size)
	if err != nil {
		b.Fatal(err)
	}
	res := process(t)
	sfname := fileName("benchmarkstats", size, size)
	b.ResetTimer()
	for n := 0; n < b.N; n++ {
		writeToFile(sfname, res)
	}
}

func Benchmark_all(b *testing.B) {
	dfname := fileName("benchmark", size, size)
	for n := 0; n < b.N; n++ {
		data, err := readFromFile(dfname, size, size)
		if err != nil {
			log.Fatalln(err)
		}

		stats := process(data)
		if err != nil {
			log.Fatalln(err)
		}

		sfname := fileName("benchmarkstats", size, size)
		err = writeToFile(sfname, stats)
		if err != nil {
			log.Fatalln(err)
		}
	}
}
