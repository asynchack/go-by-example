package main

import (
	"fmt"
	"os"
	"path/filepath"
)

func main() {
	f, err := os.CreateTemp("", "sample")
	if err != nil {
		panic(err)
	}
	defer os.Remove(f.Name())

	fmt.Println("temp file name:", f.Name())

	//
	_, err = f.Write([]byte{1, 2, 3})

	//
	dname, err := os.MkdirTemp("", "sampledir")
	fmt.Println(dname)
	defer os.RemoveAll(dname)

	//
	fname := filepath.Join(dname, "file1")
	err = os.WriteFile(fname, []byte{4, 5, 6}, 0666)

}
