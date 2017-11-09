package main

// We want to sort a list of strings by length.

type List []string

var names List

// Len is simply the length of the slice.
func (l List) Len() int {
	return len(l)
}

// Less returns true if the length of the element at index i
// is less than the length of the element at index j.
func (l List) Less(i, j int) bool {
	return len(l[i]) < len(l[j])
}

// Swap does a simple slice swap.
func (l List) Swap(i, j int) {
	l[i], l[j] = l[j], l[i]
}
