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
	res := Stack[T]{}
	for s1.Len() > 0 {
		res.Push(s1.Pop())
	}
	for s2.Len() > 0 {
		res.Push(s2.Pop())
	}
	return res
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
	stack1 := NewStack[string]()
	stack1.Push("alpha")
	stack1.Push("beta")
	stack1.Push("gamma")

	fmt.Println("generic string stack:", stack1)

	fmt.Println("Pop:  ", stack1.Pop())
	fmt.Println("Generic string stack:", stack1)
	fmt.Println()

	//generic stack: int
	stack2 := Stack[int]{1, 2, 3, 4}

	fmt.Println("Generic int stack:", stack2)

	top2 := stack2.Pop()
	fmt.Println("Pop:  ", top2)
	fmt.Println("Generic int stack:", stack2)

	stack3 := Stack[int]{4, 5, 6}
	fmt.Println("Combine:", Combine(stack2, stack3))

}
