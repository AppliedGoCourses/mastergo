package main

import (
	"flag"
	"log"
	"sync"
	"testing"
)

var once = sync.Once{}
var r, c int

func setup(b *testing.B) (filename string) {
	flag.Parse()
	r = *rows
	c = *cols
	filename = fileName("benchmark", r, c)
	once.Do(func() {
		err := generateIfNotExists(filename, r, c)
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
		_, err := readFromFile(fname, r, c)
		if err != nil {
			b.Fatal(err)
		}
	}
}

func Benchmark_process(b *testing.B) {
	fname := setup(b)
	t, err := readFromFile(fname, r, c)
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
	t, err := readFromFile(fname, r, c)
	if err != nil {
		b.Fatal(err)
	}
	res := process(t)
	sfname := fileName("benchmarkstats", r, c)
	b.ResetTimer()
	for n := 0; n < b.N; n++ {
		err := writeToFile(sfname, res)
		if err != nil {
			b.Fatal(err)
		}
	}
}

func Benchmark_all(b *testing.B) {
	dfname := fileName("benchmark", r, c)
	for n := 0; n < b.N; n++ {
		data, err := readFromFile(dfname, r, c)
		if err != nil {
			log.Fatalln(err)
		}

		stats := process(data)
		if err != nil {
			log.Fatalln(err)
		}

		sfname := fileName("benchmarkstats", r, c)
		err = writeToFile(sfname, stats)
		if err != nil {
			log.Fatalln(err)
		}
	}
}
