package main

import (
	"encoding/csv"
	"flag"
	"fmt"
	"io"
	"log"
	"os"
	"strconv"
	"sync"
	"time"

	"golang.org/x/sync/errgroup"

	gf "github.com/brianvoe/gofakeit"
	"github.com/pkg/errors"
)

type Row struct {
	Name  string
	Hrate []int
}

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

	file, rch, errch, err := readFromFile(dfname, *rows)
	if err != nil {
		log.Fatalln(err)
	}
	defer file.Close()
	// Check the error channel:
	go func() {
		select {
		case err := <-errch:
			log.Fatalln(err)
		}
	}()

	wch := process(rch, *rows)

	sfname := fileName("stats", *rows, *cols)

	err = writeToFile(sfname, wch)
	if err != nil {
		log.Fatalln(err)
	}

}

// readFromFile is a file handling wrapper for read().
// This way we can make read() testable with non-file data.
func readFromFile(name string, rows int) (*os.File, chan Row, chan error, error) {
	f, err := os.Open(name)
	if err != nil {
		return nil, nil, nil, errors.Wrapf(err, "Cannot read %s", name)
	}

	rch, ech := read(f, rows)
	return f, rch, ech, nil
}

func read(r io.Reader, rows int) (chan Row, chan error) {
	// A CSV reader is aware of the structure and syntax of a CSV file.
	// NewReader expects an io.Reader, and os.File implements io.Reader,
	// so we can simply pass in the open file.
	cr := csv.NewReader(r)

	// We do not use the current line after reading its fields,
	// so we can reuse the allocated record for performance.
	cr.ReuseRecord = true

	// Create a buffered channel for sending rows to process()
	ch := make(chan Row, rows)

	// Our goroutines might produce errors. To handle these,
	// we use an errgroup.

	var g errgroup.Group

	g.Go(func() error {
		defer close(ch)
		for {
			// Read the CSV data line by line.
			line, err := cr.Read()
			if err == io.EOF {
				// We're done.
				return nil
			}
			if err != nil {
				return errors.Wrap(err, "Cannot read row")
			}
			if len(line) == 0 { // skip empty lines
				continue
			}

			// Create a new Row and fill it with CSV data.
			cols := len(line) - 1 // do not count the name column
			row := Row{
				Name:  line[0],
				Hrate: make([]int, cols, cols),
			}
			for j := 1; j < cols; j++ {
				// All CSV data is of type string, but we want to store heart rates as integers.
				hr, err := strconv.Atoi(line[j])
				if err != nil {
					return errors.Wrapf(err, "Cannot convert string '%s' from column %d to int", line[j], j)
				}
				row.Hrate[j-1] = hr
			}
			ch <- row
		}
		return nil
	})

	// Now we must wait for all goroutines to finish, in order
	// to collect any error messages.
	// However, we also need to return the channel immediately
	// so that the data can flow to the processing stage.
	// Hence we need to wait within a goroutine, and use a
	// channel to return any error.
	errch := make(chan error)
	go func() {
		err := g.Wait()
		if err != nil {
			errch <- err
		}
	}()
	return ch, errch
}

func process(rch chan Row, rows int) chan Row {

	wch := make(chan Row, rows)

	wg := sync.WaitGroup{}

	for row, ok := <-rch; ok; row, ok = <-rch {
		wg.Add(1)
		go func(r Row) {
			wch <- simulateSlowServer(r)
			wg.Done()
		}(row)
	}
	go func() {
		wg.Wait()
		close(wch)
	}()

	return wch
}

func writeToFile(name string, ch chan Row) (err error) {
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

	return write(ch, f)
}

func write(ch chan Row, w io.Writer) error {
	cw := csv.NewWriter(w)

	// We want a header row in our output CSV file.
	cw.Write([]string{"Name", "avg", "min", "max"})

	for data, ok := <-ch; ok; data, ok = <-ch {
		// Turn our stats into strings.
		// With more than three values, a loop might be preferable.
		row := []string{
			data.Name,
			strconv.Itoa(data.Hrate[0]),
			strconv.Itoa(data.Hrate[1]),
			strconv.Itoa(data.Hrate[2]),
		}
		if err := cw.Write(row); err != nil {
			return errors.Wrapf(err, "Cannot write row '%v'", row)
		}
	}
	cw.Flush()

	return nil
}

// *** NOTE: All functions below this point are just helper functions.
// *** No need to optimize anything here.

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
			fmt.Fprintf(f, "\"%d\"", gf.Number(min, max))
			if j < cols-1 {
				fmt.Fprintf(f, ",")
			}
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
