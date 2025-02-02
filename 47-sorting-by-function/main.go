package main

import (
	"cmp"
	"fmt"
	"slices"
)

func main() {
	fruits := []string{"apple", "orange", "kiwi"}
	lenCmp := func(a, b string) int {
		return cmp.Compare(len(a), len(b))
	}

	slices.SortFunc(fruits, lenCmp)
	fmt.Println(fruits)

	type person struct {
		name string
		age  int
	}
	persons := []person{
		person{name: "Jax", age: 37},
		person{name: "TJ", age: 25},
		person{name: "Alex", age: 72},
	}

	slices.SortFunc(persons, func(a, b person) int {
		return cmp.Compare(a.age, b.age)
	})
	fmt.Println(persons)
}
