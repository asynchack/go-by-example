package main

import (
	"fmt"
	"slices"
)

func main() {

	var s []string
	fmt.Println("result: ", s, s == nil, len(s) == 0)

	s = make([]string, 3)
	fmt.Println(s, len(s), cap(s))

	//
	s[0] = "a"
	s[1] = "b"
	s[2] = "c"
	fmt.Println(s)
	fmt.Println("get:", s[2])

	s = append(s, "d")
	fmt.Println(s, len(s), cap(s))

	//
	c := make([]string, len(s))
	copy(c, s)
	fmt.Println(c, len(s), cap(s))

	//
	fmt.Println(s)
	l := s[2:5]
	fmt.Println(l, len(l), cap(l))

	//
	t := []string{"g", "h", "i"}
	fmt.Println(t)

	t2 := []string{"g", "h", "i"}
	if slices.Equal(t, t2) {
		fmt.Println("equal")
	}

	//
	twoD := make([][]int, 3)
	for i := 0; i < 3; i++ {
		temp := i + 1
		twoD[i] = make([]int, temp)
		for j := 0; j < temp; j++ {
			twoD[i][j] = j
		}
	}
	fmt.Println(twoD)
}
