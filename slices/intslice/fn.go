package intslice

type BoolComped interface {
	And(bfn BoolFn) BoolFn
	AndNot(bfn BoolFn) BoolFn
	Or(bfn BoolFn) BoolFn
}

type IntComped interface {
	And(ifn IntFn) IntFn
}

type BoolFunked interface {
	Filter([]int) []int
	Some([]int) bool
	Every([]int) bool
	BoolComped
}

type IntFunked interface {
	Map([]int) []int
	IntComped
}

type BoolFn func(int, int, []int) bool
type IntFn func(int, int, []int) int
type SortFn func(int, int) int

// Filter iterates ints and applies BoolFn on each iteration value.
// Returns the filtered list
func (bfn BoolFn) Filter(ints []int) []int {
	filtered := make([]int, 0, len(ints))

	for iter, i := range ints {
		if bfn(i, iter, ints) {
			filtered = append(filtered, i)
		}
	}

	return filtered
}

// Map iterates ints and applies IntFn to each iteration value
// Returns the mapped list
func (ifn IntFn) Map(ints []int) []int {
	length := len(ints)
	mapped := make([]int, length, length)

	for i := 0; i < length; i++ {
		mapped[i] = ifn(ints[i], i, ints)
	}

	return mapped
}

// Some iterates ints and applies BoolFn to each value
// Returns on first true case
func (bfn BoolFn) Some(ints []int) bool {
	for iter, i := range ints {
		if bfn(i, iter, ints) {
			return true
		}
	}

	return false
}

// Every iterates ints and applies BoolFn to each value
// Returns on first false case
func (bfn BoolFn) Every(ints []int) bool {
	for iter, i := range ints {
		if !bfn(i, iter, ints) {
			return false
		}
	}

	return true
}

// Find returns a pointer the value of the first match
// If not found, returns nil
func (bfn BoolFn) Find(ints []int) *int {
	for iter, i := range ints {
		if bfn(i, iter, ints) {
			return &i
		}
	}

	return nil
}

// FindIndex returns a pointer to the index of the first match
// If not found, returns nil
func (bfn BoolFn) FindIndex(ints []int) *int {
	for iter, i := range ints {
		if bfn(i, iter, ints) {
			return &iter
		}
	}

	return nil
}

// And composes two IntFn together into a new IntFn
func (ifn IntFn) And(next IntFn) IntFn {
	return func(i int, iter int, list []int) int {
		return next(ifn(i, iter, list), iter, list)
	}
}

// And composes two BoolFn together into a new BoolFn
func (bfn BoolFn) And(next BoolFn) BoolFn {
	return func(i int, iter int, list []int) bool {
		if bfn(i, iter, list) {
			return next(i, iter, list)
		}

		return false
	}
}

// AndNot composes two BoolFn into a new BoolFn, where `next` is expected to return false
func (bfn BoolFn) AndNot(next BoolFn) BoolFn {
	return func(i int, iter int, list []int) bool {
		if bfn(i, iter, list) {
			return !next(i, iter, list)
		}

		return false
	}
}

// Or composes two BoolFn into a new BoolFn, where `bfn` or `next` are expected to return true
func (bfn BoolFn) Or(next BoolFn) BoolFn {
	return func(i int, iter int, list []int) bool {
		return bfn(i, iter, list) || next(i, iter, list)
	}
}
