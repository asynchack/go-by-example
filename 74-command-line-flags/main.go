package main

import (
	"flag"
	"fmt"
)

func main() {
	wordPtr := flag.String("word", "foo", "a string")
	numPtr := flag.Int("num", 100, "a int")
	forkPtr := flag.Bool("for", true, "a bool")

	var svar string
	flag.StringVar(&svar, "svar", "bar", "a string var")

	flag.Parse()

	fmt.Println(*wordPtr)
	fmt.Println(*numPtr)
	fmt.Println(*forkPtr)
	fmt.Println(svar)

	fmt.Println(flag.Args())

}
