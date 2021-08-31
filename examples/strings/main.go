package main

import (
	"fmt"
	"strings"

	"github.com/jwhittle933/funked/slices/stringslice"
)

func main() {
	strs := []string{"one", "two", "three", "four", "five"}

	fmt.Println(stringslice.From(strs).Map(doubleWithDashes))
	fmt.Println(stringslice.StringFn(doubleWithDashes).Map(strs))
	fmt.Println(stringslice.Map(strs, doubleWithDashes))

	fmt.Println(
		stringslice.From(strs).
			Map(doubleWithDashes).
			Filter(containsDashes).
			Some(hasLength(10)),
	)

	fmt.Println(
		stringslice.Map(strs, doubleWithDashes).
			Filter(containsDashes).
			Some(hasLength(10)),
	)

	mapped := stringslice.
		StringFn(prependDashes).
		And(doubleWithDashes).
		And(appendDashes).
		Map(strs)

	fmt.Println(mapped)

	fmt.Println(
		hasLength(15).
			And(containsDashes).
			Filter(mapped),
	)
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
