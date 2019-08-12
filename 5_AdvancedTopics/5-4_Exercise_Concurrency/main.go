package main

import (
	"encoding/csv"
	"flag"
	"fmt"
	"io"
	"log"
	"os"
	"strconv"
	"time"

	gf "github.com/brianvoe/gofakeit"
	"github.com/pkg/errors"
)

type Row struct {
	Name  string
	Hrate []int
}

type Table []Row

var (
	rows  = flag.Int("rows", 1000, "Number of rows")
	cols  = flag.Int("cols", 1000, "Number of columns")
	delay = flag.Duration("delay", 1*time.Millisecond, "Delay of the simulated server's response in ms")
)

func main() {
	// By default, the code creates and uses a table of
	// 1000 rows by 1000 columns. Use the -rows and -cols
	// flags to change the table size.

	flag.Parse()

	// Generate a file name based on the # of rows and columns.
	dfname := fileName("data", *rows, *cols)

	// Generate the file only if it does not exist yet.
	generateIfNotExists(dfname, *rows, *cols)

	data, err := readFromFile(dfname, *rows, *cols)
	if err != nil {
		log.Fatalln(err)
	}

	stats := process(data)

	sfname := fileName("stats", *rows, *cols)

	err = writeToFile(sfname, stats)
	if err != nil {
		log.Fatalln(err)
	}
}

// readFromFile is a file handling wrapper for read().
// This way we can make read() testable with non-file data.
func readFromFile(name string, rows, cols int) (Table, error) {
	f, err := os.Open(name)
	if err != nil {
		return nil, errors.Wrapf(err, "Cannot read %s", name)
	}
	defer f.Close()

	return read(f, rows, cols)
}

func read(r io.Reader, rows, cols int) (Table, error) {
	// A CSV reader is aware of the structure and syntax of a CSV file.
	// NewReader expects an io.Reader, and os.File implements io.Reader,
	// so we can simply pass in the open file.
	cr := csv.NewReader(r)

	// We do not use the current line after reading its fields,
	// so we can reuse the allocated record for performance.
	cr.ReuseRecord = true

	// Pre-allocate the table structure.
	data := makeTable(rows, cols)

	for i := 0; i < rows; i++ {
		// Read the CSV data line by line.
		line, err := cr.Read()
		if err != nil {
			return nil, errors.Wrapf(err, "Cannot read row %d", i)
		}

		// Fill the current table row with name and heart rates.
		data[i].Name = line[0]
		for j := 0; j < cols; j++ {
			// All CSV data is of type string, but we want to store heart rates as integers.
			hr, err := strconv.Atoi(line[j+1])
			if err != nil {
				return nil, errors.Wrapf(err, "Cannot convert string '%s' to int", line[j])
			}
			data[i].Hrate[j] = hr
		}
	}
	return data, nil
}

func process(data Table) Table {
	rows := len(data)
	stats := makeTable(rows, 3) // We store avg, min, and max

	for i := 0; i < rows; i++ {

		stats[i] = simulateSlowServer(data[i])

	}
	return stats
}

func writeToFile(name string, t Table) (err error) {
	f, err := os.Create(name)
	if err != nil {
		return errors.Wrapf(err, "Cannot create %s", name)
	}
	defer func() {
		e := f.Close()
		if e != nil {
			err = e
		}
	}()

	return write(t, f)
}

func write(t Table, w io.Writer) error {
	cw := csv.NewWriter(w)

	// We want a header row in our output CSV file.
	cw.Write([]string{"Name", "avg", "min", "max"})

	for i := 0; i < len(t); i++ {
		// Turn our stats into strings.
		// With more than three values, a loop might be preferable.
		row := []string{
			t[i].Name,
			strconv.Itoa(t[i].Hrate[0]),
			strconv.Itoa(t[i].Hrate[1]),
			strconv.Itoa(t[i].Hrate[2]),
		}
		if err := cw.Write(row); err != nil {
			return errors.Wrapf(err, "Cannot write row %d (%v)", i, row)
		}
	}
	cw.Flush()

	return nil
}

// *** NOTE: All functions below this point are just helper functions.
// *** No need to optimize anything here.

func makeTable(r, c int) Table {
	// Return value "Table" does not need to be a pointer, since it represents
	// a slice header that consists of len, cap, and a pointer to the actual
	// data. Remember the lecture on slices!
	t := make(Table, r, r) // set len and cap to # of rows
	for i := 0; i < r; i++ {
		// Pre-allocate a row
		t[i].Hrate = make([]int, c, c) // set len and cap to # of cols
	}
	return t
}

func fileName(p string, r, c int) string {
	return fmt.Sprintf("%s%sx%s.csv", p, strconv.Itoa(r), strconv.Itoa(c))
}

func generateIfNotExists(name string, rows, cols int) error {
	_, err := os.Stat(name)
	if err == nil {
		// File exists, no need for creating one.
		return nil
	}
	if !os.IsNotExist(err) {
		// Only a "not exists" error is expected here.
		return errors.Wrap(err, "Unexpected error on os.Stat")
	}

	f, err := os.Create(name)
	if err != nil {
		return errors.Wrapf(err, "Cannot create %s", name)
	}
	defer f.Close()

	for i := 0; i < rows; i++ {
		fmt.Fprintf(f, "\"%s\",", gf.Name())
		min := gf.Number(80, 100)
		max := gf.Number(160, 180)
		for j := 0; j < cols; j++ {
			fmt.Fprintf(f, "\"%d\",", gf.Number(min, max))
		}
		fmt.Fprintln(f)
	}
	return nil
}

// This function simulates a server that stores and evaluates
// all training data. As a matter of fact, it needs some time to
// send the results back.
func simulateSlowServer(data Row) Row {
	// simulate work
	time.Sleep(*delay)

	sum := 0   // used for calculating average heard frequency
	min := 999 // larger than any possible human heart rate
	max := 0
	cols := len(data.Hrate)

	for j := 0; j < cols; j++ {
		hr := data.Hrate[j]
		sum += hr
		if hr < min {
			min = hr
		}
		if hr > max {
			max = hr
		}
	}
	stats := Row{
		Name: data.Name,
		Hrate: []int{
			sum / cols,
			min,
			max,
		},
	}
	return stats
}
