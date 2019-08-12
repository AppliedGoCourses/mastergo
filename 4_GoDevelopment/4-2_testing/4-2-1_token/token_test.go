package token

import (
	"reflect"
	"testing"
)

func TestScan(t *testing.T) {
	input := "  Always  test     your code  "
	expected := []string{"Always", "test", "your", "code"}

	// Here we call the function to be tested.
	actual := Scan(input)

	// Now we examine the output.

	if !reflect.DeepEqual(actual, expected) {
		t.Errorf("Expected: %#v\nActual:   %#v\n", expected, actual)
	}
}
