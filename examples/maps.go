package main

import (
	"fmt"
	"maps"
)

func main() {
	m := make(map[string]int)
	m["k1"] = 10
	m["k2"] = 100

	fmt.Println("map:", m)

	v1 := m["k1"]
	fmt.Println(v1)

	fmt.Println(len(m))

	delete(m, "k2")
	fmt.Println(len(m))

	clear(m)
	fmt.Println(m)
	fmt.Println(len(m))

	//
	n := map[string]int{"foo": 1, "bar": 2}
	n2 := map[string]int{"foo": 1, "bar": 2}
	if maps.Equal(n, n2) {
		fmt.Println("equals")
	}
}
