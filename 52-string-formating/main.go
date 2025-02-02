package main

import (
	"fmt"
	"os"
)

type point struct {
	x, y int
}

func main() {
	p := point{x: 3, y: 4}
	fmt.Printf("s1, %v\n", p)
	fmt.Printf("s2, %+v\n", p)
	fmt.Printf("s3, %#v\n", p)

	fmt.Printf("type, %T\n", p)
	fmt.Printf("pointer, %p\n", &p)
	fmt.Printf("bool, %t\n", true)
	fmt.Printf("int, %d\n", 123)
	fmt.Printf("hex, %x\n", 123)
	fmt.Printf("bin, %b\n", 14)
	fmt.Printf("char, %c\n", 97)

	fmt.Printf("float, %f\n", 1.23)
	fmt.Printf("float2 %e\n", 1.23456789)
	fmt.Printf("float2 %E\n", 1.23456789)

	fmt.Printf("string1, %s\n", "string")
	fmt.Printf("string2, %s\n", "\"string\"")
	fmt.Printf("string2, %q\n", "\"string\"")
	fmt.Printf("hex, %x\n", "hex this")

	fmt.Printf("width, |%6d|%6d|\n", 12, 34)
	fmt.Printf("width, |%-6d|%-6d|\n", 12, 34)
	fmt.Printf("width, |%6s|%6s|\n", "aa", "bb")
	fmt.Printf("width, |%-6s|%-6s|\n", "aa", "bb")
	fmt.Printf("width, |%6.2f|%6.2f|\n", 12.2, 34.56)
	fmt.Printf("width, |%-6.2f|%-6.2f|\n", 12.2, 34.56)

	s := fmt.Sprintf("a string for : %s\n", "good")
	fmt.Println(s)

	fmt.Fprintf(os.Stderr, "second string for : %s\n", "upup")
}
