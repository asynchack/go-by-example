package main

import (
	"fmt"
	"os"
	"strings"
)

func main() {
	os.Setenv("foo", "foo")
	fmt.Println(os.Getenv("foo"))
	fmt.Println(os.Getenv("bar"))

	for _, e := range os.Environ() {
		pair := strings.SplitN(e, "=", 2)
		fmt.Println(pair[0])
	}

}
