package main

import (
	"flag"
	"log"
	"sync"
	"testing"
)

var (
	once = sync.Once{}
	r, c int
)

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
		f, out, errch, err := readFromFile(fname, r)
		if err != nil {
			b.Fatal(err)
		}
		defer f.Close()
		go func() {
			select {
			case err := <-errch:
				log.Fatalln(err)
			}
		}()
		sum := 0
		for row, ok := <-out; ok; row, ok = <-out {
			sum += row.Hrate[0]
		}
	}
}

func Benchmark_process(b *testing.B) {
	fname := setup(b)
	f, ch, _, err := readFromFile(fname, r)
	defer f.Close()
	b.ResetTimer()
	if err != nil {
		b.Fatal(err)
	}
	for n := 0; n < b.N; n++ {
		out := process(ch, r)
		sum := 0
		for row, ok := <-out; ok; row, ok = <-out {
			sum += row.Hrate[0]
		}
	}
}

func Benchmark_write(b *testing.B) {
	fname := setup(b)
	f, rch, _, err := readFromFile(fname, r)
	if err != nil {
		b.Fatal(err)
	}
	defer f.Close()
	wch := process(rch, r)
	sfname := fileName("benchmarkstats", r, c)
	b.ResetTimer()
	for n := 0; n < b.N; n++ {
		err := writeToFile(sfname, wch)
		if err != nil {
			b.Fatal(err)
		}
	}
}

func Benchmark_all(b *testing.B) {
	dfname := setup(b)
	for n := 0; n < b.N; n++ {
		f, rch, _, err := readFromFile(dfname, r)
		if err != nil {
			b.Fatal(err)
		}
		defer f.Close()

		wch := process(rch, r)
		if err != nil {
			b.Fatal(err)
		}

		sfname := fileName("benchmarkstats", r, c)
		err = writeToFile(sfname, wch)
		if err != nil {
			b.Fatal(err)
		}
	}
}
