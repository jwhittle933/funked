package main

import (
	"fmt"
	"strings"

	"github.com/jwhittle933/funked/slices/stringslice"
)

func main() {
	strings := []string{"one", "two", "three", "four", "five"}

	fmt.Println(stringslice.From(strings).Map(doubleWithDashes))
	fmt.Println(stringslice.StringFn(doubleWithDashes).Map(strings))
	fmt.Println(stringslice.Map(strings, doubleWithDashes))

	mapped := stringslice.
		StringFn(prependDashes).
		And(doubleWithDashes).
		And(appendDashes).
		Map(strings)

	fmt.Println(mapped)

	filtered := stringslice.
		BoolFn(hasLength(15)).
		And(containsDashes).
		Filter(mapped)

	fmt.Println(filtered)
}

func doubleWithDashes(s string, _ int, _ []string) string {
	return s + "--" + s
}

func appendDashes(s string, _ int, _ []string) string {
	return s + "--"
}

func prependDashes(s string, _ int, _ []string) string {
	return "--" + s
}

func hasLength(l int) stringslice.BoolFn {
	return func(s string, _ int, _ []string) bool {
		return len(s) >= l
	}
}

func containsDashes(s string, _ int, _ []string) bool {
	return strings.Contains(s, "-")
}