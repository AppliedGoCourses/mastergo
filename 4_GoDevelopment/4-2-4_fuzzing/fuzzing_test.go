package fuzzing

import (
	"testing"
	"unicode"
)

func FuzzScan(f *testing.F) {
	seedInputs := []string{
		"a b c",
		// "a   b  \t  c",
		// "\t  \n \t",
		// "abc",
		// "",
	}

	// Create seed test cases
	for _, input := range seedInputs {
		f.Add(input)
	}
	f.Fuzz(func(t *testing.T, data string) {
		fields := Scan(data)
		if !noWhitespaceLeft(fields) {
			t.Errorf("Scan(%q) = %#v; whitespace left", data, fields)
		}
		if !noEmptyFields(fields) {
			t.Errorf("Scan(%q) = %#v; empty fields", data, fields)
		}
	})

}

// Verify that no white space is left in the fields.
func noWhitespaceLeft(fields []string) bool {
	for _, field := range fields {
		for _, char := range field {
			if unicode.IsSpace(char) {
				return false
			}
		}
	}
	return true
}

// Verify that no empty fields are present
func noEmptyFields(fields []string) bool {
	if len(fields) == 0 {
		return true // the slice as a whole may be empty - no error
	}
	for _, field := range fields {
		if len(field) == 0 {
			return false
		}
	}
	return true
}
