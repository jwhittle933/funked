package intslice

import (
	"sort"
	"strconv"
)

// IntSlicedAssertion convenience interface for
// all intslice funcs that return a bool
type IntSlicedAssertion interface {
	Some(fn BoolFn) bool
	Every(fn BoolFn) bool
	Includes(integer int) bool
	Empty() bool
}

// IntSlicedItem convenience interface for
// all intslice funcs that return an item from the Slice
type IntSlicedItem interface {
	Find(fn BoolFn) *int
	FindIndex(fn BoolFn) *int
	IndexOf(integer int) *int
	First() *int
	Last() *int
	At(index int) *int
}

// IntSlicer convenience interface for all
// intslice funcs that return another Slice
type IntSlicer interface {
	Filter(fn BoolFn) Slice
	Map(fn IntFn) Slice
	Prepend(integer int) Slice
	Copy(dst Slice) Slice
	Sort() []int
}

// IntSliced convenience interface for Slice API
type IntSliced interface {
	Join(sep string) string
	Slicer
	SlicedAssertion
	SlicedItem
}

type Slice []int

func From(ints []int) Slice {
	return ints
}

func Copy(ints []int) Slice {
	c := make([]int, 0, len(ints))
	copy(c, ints)
	return From(c)
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

// Includes searches the slice for the provided int
// and returns true if found, false otherwise.
func (s Slice) Includes(integer int) bool {
	for _, i := range s {
		if integer == i {
			return true
		}
	}

	return false
}

// IndexOf search for the index of the provided integer
// and returns a pointer to the int if found, nil otherwise
func (s Slice) IndexOf(integer int) *int {
	for iter, i := range s {
		if integer == i {
			return &iter
		}
	}

	return nil
}

// Join joins each slice member by a separator, skipping
// the last iteration.
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
	if index < 0 || len(s)-1 < index {
		return nil
	}

	return &s[index]
}

func (s Slice) Prepend(i int) Slice {
	out := make([]int, 1, len(s)+1)
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
	return dst
}
