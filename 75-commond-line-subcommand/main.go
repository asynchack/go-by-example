package main

import (
	"flag"
	"fmt"
	"os"
)

func main() {
	foo := flag.NewFlagSet("foo", flag.ExitOnError)

	fooEnable := foo.Bool("enable", true, "enable or not")
	fooName := foo.String("name", "good", "your name")

	bar := flag.NewFlagSet("bar", flag.ExitOnError)
	barInt := bar.Int("level", 666, "a int")

	if len(os.Args) < 2 {
		fmt.Println("expected 'foo' or 'bar' subcommands")
		os.Exit(1)
	}

	switch os.Args[1] {
	case "foo":
		foo.Parse(os.Args[2:])
		fmt.Println(*fooEnable)
		fmt.Println(*fooName)
	case "bar":
		bar.Parse(os.Args[2:])
		fmt.Println(*barInt)
	default:
		fmt.Println("neither foo or bar subcommands")
		os.Exit(1)
	}
}
