package main

import (
	"os"
	"text/template"
)

func main() {
	t1 := template.New("t1")
	t1, err := t1.Parse("values is {{.}}\n")
	if err != nil {
		panic(err)
	}
	t1 = template.Must(t1.Parse("values is : {{.}}\n"))
	t1.Execute(os.Stdout, "good")
	t1.Execute(os.Stdout, 5)
	t1.Execute(os.Stdout, []string{"apple", "netflex"})

	Create := func(name, text string) *template.Template {
		return template.Must(template.New(name).Parse(text))
	}

	t2 := Create("t2", "Name {{.Name}}\n")
	t2.Execute(os.Stdout, struct {
		Name string
	}{Name: "wang"})

	t2.Execute(os.Stdout, map[string]string{
		"Name": "wang",
	})

	t3 := Create("t3",
		"{{if . -}} yes {{else -}} no {{end}}\n")
	t3.Execute(os.Stdout, "not empty")
	t3.Execute(os.Stdout, "")

	t4 := Create("t4", "{{ range .}}{{ . }} {{end}}\n")
	t4.Execute(os.Stdout, []string{"apple", "google", "mircosoft"})
}
