package main

import (
	"encoding/xml"
	"fmt"
)

type Plant struct {
	XMLName xml.Name `xml:"plant"`
	Id      int      `xml:"id,attr"`
	Name    string   `xml:"name"`
	Origin  []string `xml:"origin"`
}

func (p Plant) String() string {
	return fmt.Sprintf("Plant id=%v, name=%v, orgin=%v", p.Id, p.Name, p.Origin)
}

func main() {
	// marshal
	coffee := &Plant{Id: 27, Name: "coffee"}
	coffee.Origin = []string{"Ethiopia", "Brazil"}

	out, _ := xml.MarshalIndent(coffee, " ", " . ")
	fmt.Printf("%T\n", out)
	fmt.Println(string(out))

	// add header
	fmt.Println(xml.Header + string(out))

	// ummarshal
	var p Plant
	if err := xml.Unmarshal(out, &p); err != nil {
		panic(err)
	}
	fmt.Println(p)

	// nest element
	tomato := &Plant{Id: 81, Name: "tomato"}
	tomato.Origin = []string{"Mexico", "California"}

	type Nesting struct {
		XMLName xml.Name `xml:"nesting"`
		Plants  []*Plant `xml:"parent>child>plant"`
	}

	n := &Nesting{}
	n.Plants = []*Plant{coffee, tomato}

	out2, _ := xml.MarshalIndent(n, " ", "  ")
	fmt.Println(string(out2))
}
