package stringslice

// SlicedAssertion convenience interface for
// all stringslice funcs that return a bool
type SlicedAssertion interface {
	Some(Slice, BoolFn) bool
	Every(Slice, BoolFn) bool
	Includes(Slice, int) bool
	Empty(Slice) bool
}

// SlicedItem convenience interface for
// all stringslice funcs that return an item from the Slice
type SlicedItem interface {
	Find(Slice, BoolFn) *int
	FindIndex(Slice, BoolFn) *int
	IndexOf(Slice, int) *int
	First(Slice) *int
	Last(Slice) *int
	At(Slice, int) *int
}

// Slicer convenience interface for all
// stringslice funcs that return another Slice
type Slicer interface {
	Filter(Slice, BoolFn) []int
	Map(Slice, StringFn) []int
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

func Filter(s Slice, bfn BoolFn) []string {
	return s.Filter(bfn)
}

func Map(s Slice, sfn StringFn) []string {
	return s.Map(sfn)
}

func Some(s Slice, bfn BoolFn) bool {
	return s.Some(bfn)
}

func Every(s Slice, bfn BoolFn) bool {
	return s.Every(bfn)
}

func Find(s Slice, bfn BoolFn) *string {
	return s.Find(bfn)
}

func FindIndex(s Slice, bfn BoolFn) *int {
	return s.FindIndex(bfn)
}

func Includes(s Slice, str string) bool {
	return s.Includes(str)
}

func IndexOf(s Slice, str string) *int {
	return s.IndexOf(str)
}

func Join(s Slice, sep string) string {
	return s.Join(sep)
}

func Last(s Slice) *string {
	return s.Last()
}

func First(s Slice) *string {
	return s.First()
}

func At(s Slice, index int) *string {
	return s.At(index)
}

func Prepend(s Slice, str string) Slice {
	return s.Prepend(str)
}

func Sort(s Slice) Slice {
	return s.Sort()
}

func Empty(s Slice) bool {
	return s.Empty()
}
