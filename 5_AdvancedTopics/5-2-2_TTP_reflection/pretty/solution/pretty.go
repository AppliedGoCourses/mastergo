package pretty

import (
	"fmt"
	"reflect"
	"strings"
)

func Print(i interface{}) {
	p := pretty{visited: map[uintptr]struct{}{}}
	p.print(reflect.ValueOf(i), "", 0)
}

type pretty struct {
	visited map[uintptr]struct{}
}

func (p *pretty) markVisited(v reflect.Value) {
	p.visited[v.Pointer()] = struct{}{}
}

func (p *pretty) isVisited(v reflect.Value) bool {
	_, ok := p.visited[v.Pointer()]
	return ok
}

func (p *pretty) print(v reflect.Value, prefix string, depth int) {

	// Only print until reaching a maximum depth
	if depth > 10 {
		fmt.Println("!(DEPTH EXCEEDED)")
		return
	}

	// If the current value is invalid, do not evaluate it further.
	if !v.IsValid() {
		return
	}

	// This is used a couple of times below.
	indent := strings.Repeat(" ", depth*4)

	switch v.Kind() {

	case reflect.Ptr:

		// If the pointer is nil, print <nil> and return.

		if v.IsNil() {
			fmt.Printf("%s%s (*%s): <nil>\n", indent, prefix, v.Type().Elem())
			return
		}

		// If the pointer has been visited before, print a note
		// and return, in order to avoid infinite recursion.

		if p.isVisited(v) {
			fmt.Printf("%s %s (!%s ALREADY VISITED)\n", indent, prefix, v.Type())
			return
		}

		// Print "*" to indicate a pointer, then do a pointer indirection
		// via Elem(), and pretty-print the value behind the pointer.
		// We do this by recursively calling Print() with the pointer's value.
		// To avoid infinit recursion, we mark the current pointer as visited.

		p.markVisited(v)
		fmt.Printf("%s*%s (%s):\n", indent, prefix, v.Type().Elem())
		p.print(v.Elem(), "", depth+1)

	case reflect.Struct:

		// Print the type's name and an opening brace.
		// Loop over the fields and pretty-print each of them.
		// Print a closing brace.

		fmt.Printf("%s%s: {\n", indent, v.Type())
		for n := 0; n < v.NumField(); n++ {
			// Here, we pass the field names down to p.print(),
			// to have them printed with the proper indentation
			// level.
			p.print(v.Field(n), v.Type().Field(n).Name, depth+1)
		}
		fmt.Printf("%s}\n", indent)

	default:

		fmt.Printf("%s%s (%s): %v\n", indent, prefix, v.Type(), v)
	}
}
