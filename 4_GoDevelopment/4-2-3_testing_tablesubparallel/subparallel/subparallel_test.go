package subparallel

import (
	"fmt"
	"reflect"
	"testing"
)

func TestScan(t *testing.T) {
	t.Run("Subtest 1", func(t *testing.T) {
		t.Parallel()
		input := "  Test     your code  "
		expected := []string{"Test", "yer", "code"}

		fmt.Println("Subtest 1")
		actual := Scan(input)

		if !reflect.DeepEqual(actual, expected) {
			t.Errorf("Expected: %#v\nActual:   %#v\n", expected, actual)
		}
	})
	t.Run("Subtest 2", func(t *testing.T) {
		t.Parallel()
		input := "  Test     your code  "
		expected := []string{"Test", "your", "code"}

		fmt.Println("Subtest 2")
		actual := Scan(input)

		if !reflect.DeepEqual(actual, expected) {
			t.Errorf("Expected: %#v\nActual:   %#v\n", expected, actual)
		}
	})
}
