package main

import (
	"fmt"
	"reflect"
	"strings"
	"unicode/utf8"
	"unsafe"
)

func main() {

	// Unicode characters encoded in UTF-8 can have different byte lengths
	fmt.Println("len(a):", len("a"))
	fmt.Println("len(\"ä\"):", len("ä"))
	fmt.Println("len(\"走\"):", len("走"))

	fmt.Println()

	// Note how the range index advances. Only few UTF-8-encoded characters fit into one byte. The "rune" type is an alias of int32 (but TypeOf() doesn't care).
	for i, v := range "aä走." {
		fmt.Println("range:", i, v, string(v), reflect.TypeOf(v), unsafe.Sizeof(v))
	}

	// Both IndexByte and IndexRune return the byte index.
	fmt.Println("IndexByte:", strings.IndexByte("aä走.", '.'))
	fmt.Println("IndexRune:", strings.IndexRune("aä走.", '走'))

	fmt.Println("Rune count:", utf8.RuneCountInString("aä走."))

}
