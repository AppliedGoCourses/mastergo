package main

import (
	"fmt"
	"reflect"
)

type Mineral struct {
	Name   string
	Weight int `unit:"g"`
	Opaque bool
	id     int64
}

func examineType(t reflect.Type) {

	fmt.Println("\n*** Examining the type ***")

	// Reflect upon the type of t:

	fmt.Println("Type of t:", t)
	fmt.Println("Kind of t:", t.Kind())

	fmt.Println("Is t a struct?", t.Kind() == reflect.Struct)

	//
	isPtr := t.Kind() == reflect.Ptr
	fmt.Println("Is t a pointer?", isPtr)

	if isPtr {
		fmt.Println("Changing t to t.Elem()")
		t = t.Elem()
		fmt.Println("Is t now a struct?", t.Kind() == reflect.Struct)

	}

	// Depending on the actual type, reflect.Type provides
	// methods for inspecting that type further.
	// In case of a struct, we can examine its fields:

	field, ok := t.FieldByName("Weight")
	if !ok {
		fmt.Println("Field 'Weight' does not exist")
	}
	fmt.Println("Field name:", field.Name)
	fmt.Println("Field type:", field.Type)

	// field is a reflect.StructField and has a field named "Tag"
	// that represents the field tag. The Get() method
	tag := field.Tag
	fmt.Println("Field tag:", tag)
	unit := tag.Get("unit")
	fmt.Println("Field tag 'unit' contains:", unit)
}

func examineValue(v reflect.Value) {

	fmt.Println("\n*** Examining the value ***")

	fmt.Println("Value of v:", v)
	fmt.Println("Type of v:", v.Type())
	fmt.Println("Kind of v:", v.Kind())

	_, ok := v.Interface().(Mineral)
	if ok {
		field := v.FieldByName("Weight")
		fmt.Println("Value of field 'Weight':", field)
	}
}

func modifyValue(v reflect.Value, weight int) {

	fmt.Println("\n*** Modifying the value ***")

	// To change a reflection value, the value must
	// a) be addressable, and
	// b) be exported.
	// Addressable means that changing the value would change the
	// original value.

	fmt.Println("Is v addressable?", v.CanAddr())
	fmt.Println("Can we modify v?", v.CanSet())

	// We cannot get the field from the pointer (through v.FieldByName(...),
	// and (*v).FieldByName() does not work either, as it triggers an error:
	// "invalid indirect of v (type reflect.Value)"
	//
	// Method Elem() does the required pointer indirection:
	field := v.Elem().FieldByName("Weight")
	fmt.Println("Can we modify the field 'Weight'?", field.CanSet())

	// We stop here if v cannot be modified.
	if !field.CanSet() {
		return
	}

	elem := v.Elem()
	fmt.Println("Type of v.Elem():", elem.Type())
	fmt.Println("Is v.Elem() addressable?", elem.CanAddr())
	fmt.Println("Can we modify v.Elem()?", elem.CanSet())

	// Now we can fetch the field and change its value:
	field = elem.FieldByName("Weight")
	fmt.Println("Can we modify the field 'Weight' now?", field.CanSet())

	// In main() we will see whether the following SetInt
	// has an effect on the original value:
	field.SetInt(int64(weight))
}

func examineAndModify(iface interface{}, weight int) {
	examineType(reflect.TypeOf(iface))
	examineValue(reflect.ValueOf(iface))
	modifyValue(reflect.ValueOf(iface), weight)
}

func main() {

	m := Mineral{
		Weight: 3,
		Name:   "a test structure",
		Opaque: true,
	}

	fmt.Println(m)

	fmt.Println("\n\n***** Inspecting a struct *****")
	examineType(reflect.TypeOf(m))
	examineValue(reflect.ValueOf(m))
	modifyValue(reflect.ValueOf(&m), 4)

	fmt.Println("value of m after first modifyValue():", m)

	fmt.Println("\n\n***** Inspecting an interface *****")
	examineAndModify(&m, 5)

	fmt.Println("value of m after second modifyValue():", m)
}
