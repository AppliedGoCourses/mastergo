package main

import "fmt"

// The classic way: a stack of ints
type IntStack []int

func NewIntStack() *IntStack {
	return &IntStack{}
}

func (s *IntStack) Push(e int) {
	*s = append(*s, e)
}

func (s *IntStack) Pop() int {
	l := len(*s)
	if l == 0 {
		panic("empty stack")
	}
	e := (*s)[l-1]
	*s = (*s)[:l-1]
	return e
}

func (s IntStack) Len() int {
	return len(s)
}

// The generic way: a stack of elements of arbitrary type
type Stack[T any] []T

func NewStack[T any]() *Stack[T] {
	return &Stack[T]{}
}

func (s *Stack[T]) Push(e T) {
	*s = append(*s, e)
}

func (s *Stack[T]) Pop() T {
	l := len(*s)
	if l == 0 {
		panic("empty stack")
	}
	e := (*s)[l-1]
	*s = (*s)[:l-1]
	return e
}

func (s Stack[T]) Len() int {
	return len(s)
}

func Combine[T any](s1, s2 Stack[T]) Stack[T] {
	// The underlying type of Stack[T] is a slice of T,
	// and so we can use slice operations on a Stack[T]
	return append(s1, s2...)
}

func main() {

	// classic int stack
	is := NewIntStack()
	is.Push(1)
	is.Push(2)
	fmt.Println("classic int stack:", is)
	fmt.Println("Pop:", is.Pop())
	fmt.Println("classic int stack:", is)
	fmt.Println()

	// generic stack: string
	words := NewStack[string]()
	words.Push("alpha")
	words.Push("beta")
	words.Push("gamma")

	fmt.Println("generic string stack:", words)

	fmt.Println("Pop:  ", words.Pop())
	fmt.Println("Generic string stack:", words)
	fmt.Println()

	//generic stack: int
	numbers := Stack[int]{1, 2, 3, 4}

	fmt.Println("Generic int stack:", numbers)

	top := numbers.Pop()
	fmt.Println("Pop:  ", top)
	fmt.Println("Generic int stack:", numbers)

	moreNumbers := Stack[int]{4, 5, 6}
	fmt.Println("Combine:", Combine(numbers, moreNumbers))

}
