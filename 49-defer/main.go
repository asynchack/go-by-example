package main

import (
	"fmt"
	"os"
)

func createFile(name string) *os.File {
	fmt.Println("creating")
	f, err := os.Create(name)
	if err != nil {
		panic(err)
	}
	return f
}

func writeFile(f *os.File) {
	fmt.Fprintln(f, "data")
}

func closeFile(f *os.File) {
	fmt.Println("closing")
	err := f.Close()
	if err != nil {
		panic(err)
	}
}
func main() {
	f := createFile("/tmp/defer.txt")
	defer closeFile(f)
	writeFile(f)
}
