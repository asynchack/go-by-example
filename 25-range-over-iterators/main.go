package main

import (
	"fmt"
	"iter"
	"slices"
)

func SliceIndex[S ~[]E, E comparable](s S, v E) int {
	for i := range s {
		if v == s[i] {
			return i
		}
	}

	return -1
}

type List[T any] struct {
	head, tail *element[T]
}

type element[T any] struct {
	next *element[T]
	val  T
}

func (lst *List[T]) Push(v T) {
	if lst.head == nil {
		lst.head = &element[T]{val: v}
		lst.tail = lst.head
	} else {
		lst.tail.next = &element[T]{val: v}
		lst.tail = lst.tail.next
	}
}

func (lst *List[T]) AllElements() []T {
	var eles []T
	for e := lst.head; e != nil; e = e.next {
		eles = append(eles, e.val)
	}
	return eles
}

// for range iterator
func (lst *List[T]) All() iter.Seq[T] {
	return func(yield func(T) bool) {
		for e := lst.head; e != nil; e = e.next {
			if !yield(e.val) {
				return
			}
		}
	}
}

// can be infinite

func genFib() iter.Seq[int] {
	return func(yield func(int) bool) {
		a, b := 1, 1

		for {
			if !yield(a) {
				return
			}
			a, b = b, a+b
		}

	}
}
func main() {
	lst := List[int]{}
	lst.Push(1)
	lst.Push(2)
	lst.Push(3)

	for i := range lst.All() {
		fmt.Println(i)
	}

	res := slices.Collect(lst.All())
	fmt.Println(res)

	for i := range genFib() {
		if i >= 10 {
			break
		}
		fmt.Println(i)
	}
}
