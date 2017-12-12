package main

import (
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

func Benchmark_read(b *testing.B) {
	fname := setup(b)
	b.ResetTimer()
	for n := 0; n < b.N; n++ {
		f, _, _, err := readFromFile(fname, size)
		if err != nil {
			b.Fatal(err)
		}
		f.Close()
	}
}

func Benchmark_process(b *testing.B) {
	fname := setup(b)
	f, ch, _, err := readFromFile(fname, size)
	defer f.Close()
	b.ResetTimer()
	if err != nil {
		b.Fatal(err)
	}
	for n := 0; n < b.N; n++ {
		_ = process(ch, size)
	}
}

func Benchmark_write(b *testing.B) {
	fname := setup(b)
	f, rch, _, err := readFromFile(fname, size)
	if err != nil {
		b.Fatal(err)
	}
	defer f.Close()
	wch := process(rch, size)
	sfname := fileName("benchmarkstats", size, size)
	b.ResetTimer()
	for n := 0; n < b.N; n++ {
		err := writeToFile(sfname, wch)
		if err != nil {
			b.Fatal(err)
		}
	}
}

func Benchmark_all(b *testing.B) {
	dfname := fileName("benchmark", size, size)
	for n := 0; n < b.N; n++ {
		f, rch, _, err := readFromFile(dfname, size)
		if err != nil {
			b.Fatal(err)
		}
		defer f.Close()

		wch := process(rch, size)
		if err != nil {
			b.Fatal(err)
		}

		sfname := fileName("benchmarkstats", size, size)
		err = writeToFile(sfname, wch)
		if err != nil {
			b.Fatal(err)
		}
	}
}
