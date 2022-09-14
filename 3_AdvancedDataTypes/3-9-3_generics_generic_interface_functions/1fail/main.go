package main

import "fmt"

// We want a generic method "Clone() Cloner".

type Cloner interface {
	Clone() Cloner
}

type CloneableSlice []int

func (c CloneableSlice) Clone() CloneableSlice {
	res := make(CloneableSlice, len(c))
	copy(res, c)
	return res
}

type CloneableMap map[int]int

func (c CloneableMap) Clone() CloneableMap {
	res := make(CloneableMap, len(c))
	for k, v := range c {
		res[k] = v
	}
	return res
}

func CloneAny(c Cloner) Cloner {
	return c.Clone()
}

func main() {
	s := CloneableSlice{1, 2, 3, 4}
	fmt.Println(s.Clone())
	fmt.Println(CloneAny(s))

	m := CloneableMap{1: 1, 2: 2, 3: 3, 4: 4}
	fmt.Println(m.Clone())
	fmt.Println(CloneAny(m))
}
