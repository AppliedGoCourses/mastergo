package main

import (
	"os"
	"reflect"
	"testing"

	"github.com/appliedgocourses/quotes"
)

func TestApp_createQuote(t *testing.T) {
	tests := []struct {
		name    string
		json    []byte
		wantErr bool
	}{
		{"Alfred", []byte(`{"author": "Alfred E. Neuman", "text": "What, me worry?", "source": "MAD Magazine"}`), false},
		{"Alfred again", []byte(`{"author": "Alfred E. Neuman", "text": "What, me worry?", "source": "MAD Magazine"}`), true},
	}
	db, err := quotes.Open("testdb")
	if err != nil {
		t.Fatalf("Cannot create test DB")
	}
	app := &App{db: *db}
	defer func() {
		app.db.Close()
		os.Remove("testdb")
	}()

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := app.createQuote(tt.json); (err != nil) != tt.wantErr {
				t.Errorf("App.createQuote() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestApp_getQuote(t *testing.T) {
	tests := []struct {
		name    string
		author  string
		want    *quotes.Quote
		wantErr bool
	}{
		{"Alfred", "Alfred E. Neuman", &quotes.Quote{
			Author: "Alfred E. Neuman",
			Text:   "What, me worry?",
			Source: "MAD Magazine",
		},
			false},
	}
	db, err := quotes.Open("testdb")
	if err != nil {
		t.Fatalf("Cannot create test DB")
	}
	app := &App{db: *db}
	defer func() {
		app.db.Close()
		os.Remove("testdb")
	}()
	app.db.Create(
		&quotes.Quote{
			Author: "Alfred E. Neuman",
			Text:   "What, me worry?",
			Source: "MAD Magazine",
		})

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := app.getQuote(tt.author)
			if (err != nil) != tt.wantErr {
				t.Errorf("App.getQuote() error = %s, wantErr %s", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("App.getQuote() = %#v, want %#v", got, tt.want)
			}
		})
	}
}
