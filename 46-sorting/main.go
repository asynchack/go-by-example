package main

import (
	"fmt"
	"slices"
)

func main() {
	str := []string{"a", "c", "b"}
	slices.Sort(str)
	fmt.Println(str)

	ints := []int{1, 51, 6}
	slices.Sort(ints)
	fmt.Println(ints)

	fmt.Println(slices.IsSorted(ints))
}
