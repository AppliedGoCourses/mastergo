package main

import (
	"reflect"
	"testing"
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
	app := &App{
		storage: map[string]*Quote{},
	}
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
		want    []byte
		wantErr bool
	}{
		{"Alfred", "Alfred E. Neuman", []byte(`{
  "author": "Alfred E. Neuman",
  "text": "What, me worry?",
  "source": "MAD Magazine"
}`),
			false},
	}

	app := &App{
		storage: map[string]*Quote{
			"Alfred E. Neuman": &Quote{
				Author: "Alfred E. Neuman",
				Text:   "What, me worry?",
				Source: "MAD Magazine",
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := app.getQuote(tt.author)
			if (err != nil) != tt.wantErr {
				t.Errorf("App.getQuote() error = %s, wantErr %t", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("App.getQuote() = %s, want %s", got, tt.want)
			}
		})
	}
}
