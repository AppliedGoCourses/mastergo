package main

import "fmt"

// We want a generic method "Clone() Cloner".

type Cloner interface {
	Clone() Cloner
}

type CloneableSlice []int

// Now let's define a Clone method for CloneableSlice.
func (c CloneableSlice) Clone() Cloner {
	res := make(CloneableSlice, len(c))
	copy(res, c)
	return res
}

type CloneableMap map[int]int

// Same for CloneableMap.
func (c CloneableMap) Clone() Cloner {
	res := make(CloneableMap, len(c))
	for k, v := range c {
		res[k] = v
	}
	return res
}

// Unfortunately, this does not work:
func CloneAny(c Cloner) Cloner {
	return c.Clone()
}

func main() {
	s := CloneableSlice{1, 2, 3, 4}
	// Clone method
	fmt.Println(s.Clone())
	// Clone function
	cloned := CloneAny(s)
	if cs, ok := cloned.(CloneableSlice); ok {
		fmt.Println(cs[3])
	}

	m := CloneableMap{1: 1, 2: 2, 3: 3, 4: 4}
	// Clone method
	fmt.Println(m.Clone())
	// Clone function
	cloned = CloneAny(m)
	if cm, ok := cloned.(CloneableMap); ok {
		fmt.Println(cm[3])
	}
}
