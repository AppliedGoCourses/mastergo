package main

import (
	"fmt"
	"reflect"
)

// This is the struct we want to inspect.

type Mineral struct {
	Name string `json:"name" sql:"NAME"`
	id   int64  // internal
}

func main() {
	m := Mineral{Name: "Pyrite"}

	fmt.Println("\nm:", m)

	// First, we call inspect with m as argument.
	// As we will later see, some things do not work
	// when passing by value.
	fmt.Println("\n\n***** Struct *****")
	inspect(m)

	// Now we also pass the variable as pointer.
	fmt.Println("\n\n***** Pointer *****")
	inspect(&m)

	fmt.Println("\nm:", m)
}

// inspect() represents a function you may write some day
// for handling a range of different types. The parameter
// is an empty interface{}, and we now need to find out
// what the actual type and value behind this interface are.
func inspect(intf interface{}) {

	// Step 1: We turn intf into a reflect.Type
	// and pass it to examineType().
	examineType(reflect.TypeOf(intf))

	// Step 2: We turn inft into a reflect.Value
	// and pass it to examineAndSetValue().
	examineAndSetValue(reflect.ValueOf(intf))
}

func examineType(t reflect.Type) {
	fmt.Println("\n*** Examining the type ***")

	// Printing t itself reveals its type including a package qualifier.
	fmt.Println("Type of t:", t)

	// Type.Name() returns the name of the actual type.
	fmt.Println("Name of t:", t.Name())

	// Type.Kind() returns the underlying standard type, e.g.
	// "struct" for our Mineral type.
	fmt.Println("Kind of t:", t.Kind())

	// The reflect package contains all standard types.
	// We can compare Kind() against a type.
	// In this case, we want to find out if the variable was
	// passed in as a pointer.
	isPtr := t.Kind() == reflect.Ptr
	fmt.Println("Is t a pointer?", isPtr)

	// If the value is a pointer, we want to retrieve the value
	// behind the pointer.
	// Type.Elem() does this. The "Element" of a pointer is the
	// value it points to.
	// Elem() also works for other types that have elements:
	// arrays, slices, maps, and channels.
	// As most other methods in package reflect, Elem() panics
	// if called on a type that cannot have elements.
	if isPtr {
		fmt.Println("Changing t to t.Elem()")

		// Elem() returns a reflect.Type, so we can assign the
		// result right back to t.
		t = t.Elem()

		fmt.Println("Type of t:", t)
		fmt.Println("Kind of t:", t.Kind())
		fmt.Println("Is t now a struct?", t.Kind() == reflect.Struct)
	}

	// We want to specifically examine our Mineral struct, so
	// if the current type is not of kind "struct", we exit here.
	if t.Kind() != reflect.Struct {
		return
	}

	// Type.FieldByName() works for struct types. It returns
	// the named field if it exists.
	field, ok := t.FieldByName("Name")
	if !ok {
		fmt.Println("Field not found")
		return
	}

	// We can query the field for properties like Name, Type, or Tag.
	fmt.Println("Field name:", field.Name)
	fmt.Println("Field type:", field.Type)

	// Tag is the struct field tag used by packages like encoding/json.
	tag := field.Tag
	fmt.Println("Field tag:", tag)

	// Tag.Get expects a key name and returns the corresponding value.
	jsonName := tag.Get("json")
	fmt.Println("Field tag 'json' contains:", jsonName)

}

func examineAndSetValue(v reflect.Value) {
	fmt.Println("\n*** Examining the value ***")

	// As with examineType, we want to ensure the
	// value is no pointer.
	if v.Kind() == reflect.Ptr {
		fmt.Println("Changing v to v.Elem()")
		v = v.Elem()
	}

	fmt.Println("Value of v:", v)
	fmt.Println("Type of v:", v.Type())
	fmt.Println("Kind of v:", v.Kind())

	// We want to examine the Mineral struct, so if the value
	// is not a struct, we exit here.
	if v.Kind() != reflect.Struct {
		return
	}

	// Field(index) is another way of accessing a struct field.
	name := v.Field(0)
	fmt.Println("Value of field 'Name':", name)

	// Before attempting to change a value, we should test if it
	// can be set at all.
	// A reflect.Value can be changed if
	// (1) it is addressable (that is, the original value is still
	//	   reachable via a pointer), and
	// (2) it is not accessed through an unexported struct field.
	fmt.Println("Is field Name settable?", name.CanSet())
	fmt.Println("Is internal field id settable?", v.Field(1).CanSet())

	if !name.CanSet() {
		return
	}

	// Values can be changed through methods like Set(), SetString(),
	// SetInt(), etc.
	// The method panics if it is called on the wrong type.
	name.TryRecv()
	name.SetString("Fool's Gold")
	fmt.Println("New value of field 'Name':", name)

}
