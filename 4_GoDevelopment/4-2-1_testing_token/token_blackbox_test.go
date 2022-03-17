package token_test

import (
	"reflect"
	"testing"

	token "github.com/AppliedGoCourses/mastergo/4_GoDevelopment/4-2-1_token"
)

func TestScan(t *testing.T) {
	input := "  Always  test     your code  "
	expected := []string{"Always", "test", "your", "code"}

	// Here we call the function to be tested.
	actual := token.Scan(input)

	// Now we examine the output.

	if !reflect.DeepEqual(actual, expected) {
		t.Errorf("Expected: %#v\nActual:   %#v\n", expected, actual)
	}
}
