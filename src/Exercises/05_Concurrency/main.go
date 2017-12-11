package main

import (
	"database/sql"
	"log"

	_ "github.com/mattn/go-sqlite3"
)

func main() {
	db, err := readDB()
	if err != nil {
		log.Fatalln(err)
	}

	err = process(db)
	if err != nil {
		log.Fatalln(err)
	}

	err = writeDB(db)
	if err != nil {
		log.Fatalln(err)
	}
}

func readDB() (*sql.DB, error) {
	return nil, nil
}

func process(db *sql.DB) error {
	return nil
}

func writeDB(db *sql.DB) error {
	return nil
}
