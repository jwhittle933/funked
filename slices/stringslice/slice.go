package stringslice

import (
	"sort"
	"strings"
)

type Slice []string

func From(strs []string) Slice {
	return strs
}

func (s Slice) Filter(bfn BoolFn) Slice {
	return bfn.Filter(s)
}

func (s Slice) Map(sfn StringFn) Slice {
	return sfn.Map(s)
}

func (s Slice) Some(bfn BoolFn) bool {
	return bfn.Some(s)
}

func (s Slice) Every(bfn BoolFn) bool {
	return bfn.Every(s)
}

func (s Slice) Find(bfn BoolFn) *string {
	return bfn.Find(s)
}

func (s Slice) FindIndex(bfn BoolFn) *int {
	return bfn.FindIndex(s)
}

func (s Slice) Includes(str string) bool {
	for _, item := range s {
		if str == item {
			return true
		}
	}

	return false
}

// IndexOf search for the index of the provided string
// and returns a pointer to the index if found, nil otherwise
func (s Slice) IndexOf(str string) *int {
	for iter, item := range s {
		if str == item {
			return &iter
		}
	}

	return nil
}

func (s Slice) Join(sep string) string {
	return strings.Join(s, sep)
}

func (s Slice) Last() *string {
	if s.Empty() {
		return nil
	}

	return &s[len(s)-1]
}

func (s Slice) First() *string {
	if s.Empty() {
		return nil
	}

	return &s[0]
}

func (s Slice) At(index int) *string {
	if index < 0 || len(s)-1 < index {
		return nil
	}

	return &s[index]
}

func (s Slice) Prepend(i string) Slice {
	out := make([]string, 1, len(s)+1)
	out[0] = i
	return append(out, s...)
}

// Sort sorts the string slice
func (s Slice) Sort() Slice {
	sort.Strings(s)
	return s
}

func (s Slice) Empty() bool {
	return len(s) == 0
}

func (s Slice) Copy(dst []string) Slice {
	copy(dst, s)
	return dst
}
