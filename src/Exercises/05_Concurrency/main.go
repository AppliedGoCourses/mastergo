package main

import (
	"encoding/csv"
	"flag"
	"fmt"
	"log"
	"os"
	"strconv"

	gf "github.com/brianvoe/gofakeit"
	"github.com/pkg/errors"
)

type Row struct {
	Name  string
	Hrate []int
}

type Table []Row

var (
	rows = flag.Int("rows", 1000, "Number of rows")
	cols = flag.Int("cols", 1000, "Number of columns")
)

func main() {
	// By default, the code creates and uses a table of
	// 1000 rows by 1000 columns. Use the -rows and -cols
	// flags to change the table size.

	flag.Parse()

	data, err := read()
	if err != nil {
		log.Fatalln(err)
	}

	stats := process(data)
	if err != nil {
		log.Fatalln(err)
	}

	err = write(stats)
	if err != nil {
		log.Fatalln(err)
	}
}

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

func read() (Table, error) {
	f, err := os.Open("trainingdata.csv")
	if err != nil {
		// The file might simply not exist yet, so let's generate it.
		err = generate(*rows, *cols)
		if err != nil {
			return nil, errors.Wrap(err, "") // generate() already provides a descriptive error message.
		}
		// Now let's try again
		f, err = os.Open("trainingdata.csv")
		if err != nil {
			return nil, errors.Wrap(err, "Cannot read trainingsdata.csv")
		}
	}
	defer f.Close()

	// A CSV reader is aware of the structure and syntax of a CSV file.
	// NewReader expects an io.Reader, and os.File implements io.Reader,
	// so we can simply pass in the open file.
	cr := csv.NewReader(f)

	// Pre-allocate the table structure.
	data := makeTable(*rows, *cols)

	for i := 0; i < *rows; i++ {
		// Read the CSV data line by line.
		row, err := cr.Read()
		if err != nil {
			return nil, errors.Wrapf(err, "Cannot read row %d", i)
		}

		// Fill the current table row with name and heart rates.
		data[i].Name = row[0]
		for j := 0; j < *cols; j++ {
			// All CSV data is of type string, but we want to store heart rates as integers.
			hr, err := strconv.Atoi(row[j+1])
			if err != nil {
				return nil, errors.Wrapf(err, "Cannot convert string '%s' to int", row[j])
			}
			data[i].Hrate[j] = hr
		}
	}
	return data, nil
}

func process(data Table) Table {
	stats := makeTable(*rows, 3) // We store avg, min, and max

	for i := 0; i < *rows; i++ {
		sum := 0   // used for calculating average heard frequency
		min := 999 // larger than any possible human heart rate
		max := 0
		for j := 0; j < *cols; j++ {
			hr := data[i].Hrate[j]
			sum += hr
			if hr < min {
				min = hr
			}
			if hr > max {
				max = hr
			}
		}
		stats[i].Name = data[i].Name
		stats[i].Hrate[0] = sum / *cols
		stats[i].Hrate[1] = min
		stats[i].Hrate[2] = max
	}
	return stats
}

func write(t Table) error {
	f, err := os.Create("trainingstats.csv")
	if err != nil {
		return errors.Wrap(err, "Cannot create trainingsstats.csv")
	}
	defer f.Close()

	cw := csv.NewWriter(f)

	// We want a header row in our output CSV file.
	cw.Write([]string{"Name", "avg", "min", "max"})

	for i := 0; i < *rows; i++ {
		// Turn our stats into strings.
		// With more than three values, a loop might be preferable.
		row := []string{
			t[i].Name,
			strconv.Itoa(t[i].Hrate[0]),
			strconv.Itoa(t[i].Hrate[1]),
			strconv.Itoa(t[i].Hrate[2]),
		}
		cw.Write(row)
	}

	return nil
}

func generate(rows, cols int) error {
	f, err := os.Create("trainingdata.csv")
	if err != nil {
		return errors.Wrap(err, "Cannot open trainingsdata.csv for writing")
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
