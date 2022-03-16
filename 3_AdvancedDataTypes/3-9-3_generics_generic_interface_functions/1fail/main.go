package main

import "fmt"

// We want a generic method "Clone() Cloner".

// Let's start by defining a parametrized Cloner interface.
// At this point, the Clone() function can return any type.
// We'll refine this later.
type Cloner[C any] interface {
	Clone() C
}

type CloneableSlice []int

// Now let's define a Clone method for CloneableSlice.
func (c CloneableSlice) Clone() CloneableSlice {
	res := make(CloneableSlice, len(c))
	copy(res, c)
	return res
}

type CloneableMap map[int]int

// Same for CloneableMap.
func (c CloneableMap) Clone() CloneableMap {
	res := make(CloneableMap, len(c))
	for k, v := range c {
		res[k] = v
	}
	return res
}

// Finally, a standalone func can take the type parameter
// [T Cloner[T]] to express the fact that T must implement
// the Cloner interface for itself (T).
//
// As a result, CloneAny can take any Cloner as input.
func CloneAny[T Cloner[T]](c T) T {
	return c.Clone()
}

func main() {
	s := CloneableSlice{1, 2, 3, 4}
	// Classic clone method
	fmt.Println(s.Clone())
	// Generic clone function
	fmt.Println(CloneAny(s))

	m := CloneableMap{1: 1, 2: 2, 3: 3, 4: 4}
	// Classic clone method
	fmt.Println(m.Clone())
	// Generic clone function
	fmt.Println(CloneAny(m))
}
