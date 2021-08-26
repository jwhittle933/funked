package intslice

import (
	"sort"
	"strconv"
)

type IntSlicer interface {
	Filter(fn BoolFn) Slice
	Some(fn BoolFn) bool
	Every(fn BoolFn) bool
	Map(fn IntFn) Slice
	Find(fn BoolFn) *int
	FindIndex(fn BoolFn) *int
	Includes(integer int) bool
	IndexOf(integer int) *int
	Join(sep string) string
	First() *int
	Last() *int
	At(index int) *int
	Prepend(integer int) Slice
	Empty() bool
	Sort() Slice
	Copy(dst Slice) Slice
}

type Slice []int

func From(ints []int) Slice {
	return ints
}

func (s Slice) Filter(bfn BoolFn) Slice {
	return bfn.Filter(s)
}

func (s Slice) Some(bfn BoolFn) bool {
	return bfn.Some(s)
}

func (s Slice) Every(bfn BoolFn) bool {
	return bfn.Every(s)
}

func (s Slice) Map(ifn IntFn) Slice {
	return ifn.Map(s)
}

func (s Slice) Find(bfn BoolFn) *int {
	return bfn.Find(s)
}

func (s Slice) FindIndex(bfn BoolFn) *int {
	return bfn.FindIndex(s)
}

func (s Slice) Includes(integer int) bool {
	for _, i := range s {
		if integer == i {
			return true
		}
	}

	return false
}

func (s Slice) IndexOf(integer int) *int {
	for iter, i := range s {
		if integer == i {
			return &iter
		}
	}

	return nil
}

func (s Slice) Join(sep string) string {
	var str string
	for iter, i := range s {
		if iter == len(s)-1 {
			str += strconv.Itoa(i)
			break
		}

		str = str + strconv.Itoa(i) + sep
	}

	return str
}

func (s Slice) Last() *int {
	if s.Empty() {
		return nil
	}

	return &s[len(s)-1]
}

func (s Slice) First() *int {
	if s.Empty() {
		return nil
	}

	return &s[0]
}

func (s Slice) At(index int) *int {
	if len(s)-1 < index {
		return nil
	}

	return &s[index]
}

func (s Slice) Prepend(i int) Slice {
	length := len(s)
	out := make([]int, 1, length+1)
	out[0] = i
	return append(out, s...)
}

func (s Slice) Empty() bool {
	return len(s) == 0
}

// Sort sorts the int slice in asc order
// Uses sort.Ints (quick sort algorithm)
func (s Slice) Sort() Slice {
	sort.Ints(s)
	return s
}

func (s Slice) Copy(dst []int) Slice {
	copy(dst, s)
	return s
}
