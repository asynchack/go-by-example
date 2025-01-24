package main

import "fmt"

type person struct {
	name string
	age  int
}

func newPerson(name string) *person {
	p := person{name: name}
	p.age = 666
	return &p
}
func main() {
	fmt.Println(person{"bob", 20})
	fmt.Println(person{"alex", 10})
	fmt.Println(newPerson("joe"))

	s := person{name: "foo", age: 100}
	fmt.Println(s.age)
	sp := &s
	sp.age = 200
	fmt.Println(s.age)

	dog := struct {
		name   string
		isGool bool
	}{
		"huang",
		true,
	}

	fmt.Println(dog)

}
