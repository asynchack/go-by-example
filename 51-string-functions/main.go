package main

import (
	"fmt"
	s "strings"
)

func main() {
	var p = fmt.Println

	p("contain:", s.Contains("test", "es"))
	p("count:", s.Count("test", "t"))
	p("hasPrefix", s.HasPrefix("test", "te"))
	p("hasSufix", s.HasSuffix("test", "st"))
	p("index", s.Index("test", "e"))
	p("join", s.Join([]string{"a", "b"}, "-"))
	p("repeat", s.Repeat("a", 6))
	p("replace", s.Replace("foo", "o", "0", -10))
	p("replace", s.Replace("foo", "o", "0", 1))
	p("split", s.Split("a-b-c", "-"))
	p("toLower", s.ToLower("Test"))
	p("toUpper", s.ToUpper("Test"))
}
