package main

import "fmt"

// Constraints as interfaces

// type unions: permit all signed integers as possible instances of Signed
type SignedInts interface {
	int | int8 | int16 | int32 | int64
}

func negate[T SignedInts](n T) T {
	return -n
}

// All types whose underlying type is a signed integer
type Signed interface {
	~int | ~int8 | ~int16 | ~int32 | ~int64
}

type currency int64

// also try the deposit function with the Signed constraint
func deposit[T ~int32 | ~int64](amount T) {}

// Composing constraints by interface embedding
type Unsigned interface {
	~uint | ~uint8 | ~uint16 | ~uint32 | ~uint64 | ~uintptr
}

type Integer interface {
	Signed | Unsigned
}

type Float interface {
	~float32 | ~float64
}

type Number interface {
	Integer | Float
}

type Ordered interface {
	Number | ~string
}

// Use case for the Ordered constraint:
// a sorted binary tree (but not balanced)
// with a payload of arbitrary type
type SortedTree[T Ordered, U any] struct {
	key         T
	value       U
	left, right *SortedTree[T, U]
}

// This New function is not strictly necessary, but it clearly shows
// that the zero value of a tree is nil – the empty tree.
func NewSortedTree[T Ordered, U any]() *SortedTree[T, U] {
	return nil
}

// Upsert inserts or updates an element in the tree
func (t *SortedTree[T, U]) Upsert(key T, value U) *SortedTree[T, U] {
	if t == nil {
		return &SortedTree[T, U]{key, value, nil, nil}
	}
	if key < t.key {
		t.left = t.left.Upsert(key, value)
	} else if key > t.key {
		t.right = t.right.Upsert(key, value)
	} else {
		t.value = value
	}
	return t
}

func (t *SortedTree[T, U]) Find(key T) (U, error) {
	if t == nil {
		var zeroValue U
		return zeroValue, fmt.Errorf("key %v not found", key)
	}
	if key < t.key {
		return t.left.Find(key)
	} else if key > t.key {
		return t.right.Find(key)
	} else {
		return t.value, nil
	}
}

// A use case for the built-in comparable constraint.
// A set is like a slice but can contain a given value only once.
// To allow testing whether a value is in the set, the element type must be a comparable one.

type Set[T comparable] []T

func NewSet[T comparable]() Set[T] {
	return []T{}
}

func (s *Set[T]) Exists(e T) bool {
	for _, v := range *s {
		if v == e {
			return true
		}
	}
	return false
}

func (s *Set[T]) Add(item T) {
	if !s.Exists(item) {
		*s = append(*s, item)
	}
}

func (s *Set[T]) Get(i int) T {
	return (*s)[i]
}

func main() {

	fmt.Println("Negate:", negate(42))

	var tree *SortedTree[int, string]

	tree = tree.Upsert(1, "one")
	tree = tree.Upsert(2, "two")
	tree = tree.Upsert(3, "tree")
	tree = tree.Upsert(3, "three")
	fmt.Println("Sorted tree: ")
	fmt.Println(tree.Find(2))
	fmt.Println(tree.Find(3))

	s1 := NewSet[int]()
	s1.Add(1)
	s1.Add(2)
	s1.Add(3)
	s1.Add(1)
	fmt.Println("Set of ints:", s1)

	s2 := NewSet[string]()
	s2.Add("alpha")
	s2.Add("alpha")
	s2.Add("beta")
	fmt.Println("Set of strings:", s2)

}
