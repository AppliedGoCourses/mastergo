package main

import (
	"log"
)

type Row struct {
	Name  string
	Hfreq []int
}

type Table []Row

func main() {
	table, err := read()
	if err != nil {
		log.Fatalln(err)
	}

	err = process(table)
	if err != nil {
		log.Fatalln(err)
	}

	err = write(table)
	if err != nil {
		log.Fatalln(err)
	}
}

func read() (*Table, error) {

	return nil, nil
}

func process(t *Table) error {
	return nil
}

func write(t *Table) error {
	return nil
}
