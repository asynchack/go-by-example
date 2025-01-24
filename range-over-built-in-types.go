package main

import (
	"fmt"
)

func main() {
	nums := []int{2, 3, 4}
	total := 0
	for _, i := range nums {
		total += i
	}
	fmt.Println(total)

	for index, i := range nums {
		if index == 2 {
			fmt.Println(i)
		}
	}
	// range over maps
	kvs := map[string]string{"a": "apple", "b": "banana"}
	for k, v := range kvs {
		fmt.Printf("%s -> %s\n", k, v)

	}

	// range over key of maps
	for k := range kvs {
		fmt.Println(k)
	}

	// string and rune: return:<starting byte index of the rune and the second the rune itself>
	// 说明字母一个byte能表示，中文3个byte表示
	// 0 97
	// 1 98
	// 2 99
	// 3 65
	// 4 66
	// 5 67
	// 6 32534
	// 9 31243
	for i, c := range "abcABC编程" {
		fmt.Println(i, c)
	}
}
