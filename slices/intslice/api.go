package intslice

// SlicedAssertion convenience interface for
// all intslice funcs that return a bool
type SlicedAssertion interface {
	Some(Slice, BoolFn) bool
	Every(Slice, BoolFn) bool
	Includes(Slice, int) bool
	Empty(Slice) bool
}

// SlicedItem convenience interface for
// all intslice funcs that return an item from the Slice
type SlicedItem interface {
	Find(Slice, BoolFn) *int
	FindIndex(Slice, BoolFn) *int
	IndexOf(Slice, int) *int
	First(Slice) *int
	Last(Slice) *int
	At(Slice, int) *int
}

// Slicer convenience interface for all
// intslice funcs that return another Slice
type Slicer interface {
	Filter(Slice, BoolFn) []int
	Map(Slice, IntFn) []int
	Prepend(Slice, int) []int
	Sort(Slice) []int
}

// Sliced convenience interface for Bare API
type Sliced interface {
	Join(Slice, string) string
	Slicer
	SlicedAssertion
	SlicedItem
}

func Filter(s Slice, bfn BoolFn) Slice {
	return s.Filter(bfn)
}

func Map(s Slice, ifn IntFn) Slice {
	return s.Map(ifn)
}

func Some(s Slice, bfn BoolFn) bool {
	return s.Some(bfn)
}

func Every(s Slice, bfn BoolFn) bool {
	return s.Every(bfn)
}

func Find(s Slice, bfn BoolFn) *int {
	return s.Find(bfn)
}

func FindIndex(s Slice, bfn BoolFn) *int {
	return s.FindIndex(bfn)
}

func Includes(s Slice, includes int) bool {
	return s.Includes(includes)
}

func IndexOf(s Slice, integer int) *int {
	return s.IndexOf(integer)
}

func Join(s Slice, sep string) string {
	return s.Join(sep)
}

func Last(s Slice) *int {
	return s.Last()
}

func First(s Slice) *int {
	return s.First()
}

func Prepend(s Slice, i int) Slice {
	return s.Prepend(i)
}

func At(s Slice, index int) *int {
	return s.At(index)
}

func Empty(s Slice) bool {
	return s.Empty()
}

func Sort(s Slice) Slice {
	return s.Sort()
}
