package main

import (
	"fmt"
)

type base struct {
	num int
}

type container struct {
	base
	name string
}

type describer interface {
	describe() string
}

func (b base) describe() string {
	return fmt.Sprintf("base with num=%v\n", b.num)
}

func main() {
	co := container{
		base: base{
			num: 1,
		},
		name: "nb",
	}

	fmt.Println(co.name, co.base.num, co.num)

	var d describer
	d = co
	fmt.Println(d.describe())
}
