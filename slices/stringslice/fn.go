package stringslice

import "github.com/jwhittle933/funked/boolean"

type BoolComped interface {
	And(bfn BoolFn) BoolFn
	AndNot(bfn BoolFn) BoolFn
	Or(bfn BoolFn) BoolFn
}

type StringComped interface {
	And(sfn StringFn) StringFn
}

type BoolFunked interface {
	Filter([]string) []string
	Some([]string) bool
	Every([]string) bool
	BoolComped
}

type StringFunked interface {
	Map([]string) []string
	StringComped
}

type BoolFn func(string, int, []string) bool
type StringFn func(string, int, []string) string

func (bfn BoolFn) Filter(strs []string) []string {
	filtered := make([]string, 0, len(strs))

	for i, str := range strs {
		if bfn(str, i, strs) {
			filtered = append(filtered, str)
		}
	}

	return filtered
}

func (sfn StringFn) Map(strs []string) []string {
	length := len(strs)
	mapped := make([]string, length, length)

	for i := 0; i < length; i++ {
		mapped[i] = sfn(strs[i], i, strs)
	}

	return mapped
}

func (bfn BoolFn) Some(strs []string) bool {
	for i, str := range strs {
		if bfn(str, i, strs) {
			return true
		}
	}

	return false
}

func (bfn BoolFn) Every(strs []string) bool {
	for i, str := range strs {
		if !bfn(str, i, strs) {
			return false
		}
	}

	return true
}


// Find returns a pointer the value of the first match
// If not found, returns nil
func (bfn BoolFn) Find(strs []string) *string {
	for iter, s := range strs {
		if bfn(s, iter, strs) {
			return &s
		}
	}

	return nil
}

// FindIndex returns a pointer the value of the first match
// If not found, returns nil
func (bfn BoolFn) FindIndex(strs []string) *int {
	for iter, s := range strs {
		if bfn(s, iter, strs) {
			return &iter
		}
	}

	return nil
}


// And composes two StringFn together into a new StringFn
// Each StringFn in composition is applied to each item in the
// collection before collection iteration continues
// Example:
//  for [x, y, z] -> [first(second(third(x))), first(second(third(y))), first(second(third(z)))]
func (sfn StringFn) And(next StringFn) StringFn {
	return func(s string, iter int, strs []string) string {
		return sfn(next(s, iter, strs), iter, strs)
	}
}

// And composes two BoolFn together into a new BoolFn
func (bfn BoolFn) And(next BoolFn) BoolFn {
	return func(s string, iter int, strs []string) bool {
		return boolean.And(bfn(s, iter, strs), next(s, iter, strs))
	}
}

// AndNot composes two BoolFn into a new BoolFn, where `next` is expected to return false
func (bfn BoolFn) AndNot(next BoolFn) BoolFn {
	return func(s string, iter int, strs []string) bool {
		return boolean.AndNot(bfn(s, iter, strs), next(s, iter, strs))
	}
}

// Or composes two BoolFn into a new BoolFn, where `bfn` or `next` are expected to return true
func (bfn BoolFn) Or(next BoolFn) BoolFn {
	return func(s string, iter int, list []string) bool {
		return boolean.Or(bfn(s, iter, list), next(s, iter, list))
	}
}
