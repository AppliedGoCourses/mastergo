package quotes

import (
	"reflect"
	"testing"
)

func TestQuote_SerializeDeserialize(t *testing.T) {
	tests := []struct {
		name  string
		quote Quote
	}{
		{"01", Quote{Author: "Test", Text: "This is a test", Source: "unknown"}},
		{"02", Quote{Author: "Test", Text: "This is a test", Source: ""}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			serialized, err := tt.quote.Serialize()
			if err != nil {
				t.Errorf("Quote.Serialize() error = %v", err)
				return
			}
			restored := Quote{}
			err = restored.Deserialize(serialized)
			if err != nil {
				t.Errorf("Quote.Deserialize() error = %v", err)
				return
			}
			if !reflect.DeepEqual(tt.quote, restored) {
				t.Errorf("Quote.Serialize() -> Quote.Deserialize() = %#v, want %#v", restored, tt.quote)
			}
		})
	}
}

func TestQuote_String(t *testing.T) {
	tests := []struct {
		name  string
		quote Quote
		want  string
	}{
		{
			"StringWithSource", Quote{"Author", "Text", "Source"}, "\"Text\"\n\n(Author, Source)\n",
		},
		{
			"StringWithoutSource", Quote{"Author", "Text", ""}, "\"Text\"\n\n(Author)\n",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.quote.String(); got != tt.want {
				t.Errorf("Quote.String() = %v, want %v", got, tt.want)
			}
		})
	}
}
