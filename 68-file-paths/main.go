package main

import (
	"fmt"
	"path/filepath"
	"strings"
)

func main() {
	p := filepath.Join("dir1", "dir2", "filename")
	fmt.Println(p)

	fmt.Println(filepath.Join("dir1///", "filename"))
	fmt.Println(filepath.Join("dir1/../dir1", "filename"))

	fmt.Println(filepath.Dir(p))
	fmt.Println(filepath.Base(p))

	fmt.Println(filepath.IsAbs(p))
	fmt.Println(filepath.IsAbs(filepath.Join("/", "a", "b")))

	//
	fileName := "config.json"
	fmt.Println(filepath.Ext(fileName))

	fmt.Println(strings.TrimSuffix(fileName, filepath.Ext(fileName)))

	//
	rel, err := filepath.Rel("a/b", "a/b/c/d")
	if err != nil {
		panic(err)
	}

	fmt.Println(rel)

	rel, err = filepath.Rel("a/b", "a/x/y/z")
	if err != nil {
		panic(err)
	}
	fmt.Println(rel)
}
