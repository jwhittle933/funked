package slices

import (
	"github.com/jwhittle933/gonads/option"

	"github.com/jwhittle933/funked/boolean"
)

type BoolFn[T any] func(T, int, []T) bool
type TFn[T any] func(T, int, []T) T

type Applicator[T any] interface {
	~func(T, int, []T) bool | ~func(T, int, []T) T
}

type BoolComposer[T any] interface {
	And(bfn BoolFn[T]) BoolFn[T]
	AndNot(bfn BoolFn[T]) BoolFn[T]
	Or(bfn BoolFn[T]) BoolFn[T]
}

type TComposer[T any] interface {
	And(sfn TFn[T]) TFn[T]
}

type BoolFunked[T any] interface {
	Filter([]T) []T
	Some([]T) bool
	Every([]T) bool
	BoolComposer[T]
}

type TFunked[T any] interface {
	Map([]T) []T
	TComposer[T]
}

func (fn BoolFn[T]) Filter(list []T) []T {
	filtered := make([]T, 0, len(list))

	for i, str := range list {
		if fn(str, i, list) {
			filtered = append(filtered, str)
		}
	}

	return filtered
}

func (fn TFn[T]) Map(list []T) []T {
	length := len(list)
	mapped := make([]T, length)

	for i := 0; i < length; i++ {
		mapped[i] = fn(list[i], i, list)
	}

	return mapped
}

func (fn BoolFn[T]) Some(list []T) bool {
	for i, str := range list {
		if fn(str, i, list) {
			return true
		}
	}

	return false
}

func (fn BoolFn[T]) Every(list []T) bool {
	for i, str := range list {
		if !fn(str, i, list) {
			return false
		}
	}

	return true
}

// Find returns a pointer the value of the first match
// If not found, returns nil
func (fn BoolFn[T]) Find(list []T) option.Option[T] {
	for iter, s := range list {
		if fn(s, iter, list) {
			return option.Some(s)
		}
	}

	return option.None[T]()
}

// FindIndex returns a pointer the value of the first match
// If not found, returns nil
func (fn BoolFn[T]) FindIndex(list []T) option.Option[int] {
	for iter, s := range list {
		if fn(s, iter, list) {
			return option.Some(iter)
		}
	}

	return option.None(0)
}

// And composes two TFn's together into a new TFn.
// Each TFn in composition is applied to each item in the
// collection before collection iteration continues
// Example:
//  for [x, y, z] -> [third(second(first(x))), third(second(first(y))), third(second(first(z)))]
func (fn TFn[T]) And(next TFn[T]) TFn[T] {
	return func(s T, iter int, strs []T) T {
		return fn(next(s, iter, strs), iter, strs)
	}
}

// And composes two BoolFn together into a new BoolFn
func (fn BoolFn[T]) And(next BoolFn[T]) BoolFn[T] {
	return func(s T, iter int, strs []T) bool {
		return boolean.And(fn(s, iter, strs), next(s, iter, strs))
	}
}

// AndNot composes two BoolFn into a new BoolFn, where `next` is expected to return false
func (fn BoolFn[T]) AndNot(next BoolFn[T]) BoolFn[T] {
	return func(s T, iter int, strs []T) bool {
		return boolean.AndNot(fn(s, iter, strs), next(s, iter, strs))
	}
}

// Or composes two BoolFn into a new BoolFn, where `bfn` or `next` are expected to return true
func (fn BoolFn[T]) Or(next BoolFn[T]) BoolFn[T] {
	return func(s T, iter int, list []T) bool {
		return boolean.Or(fn(s, iter, list), next(s, iter, list))
	}
}

// StringEquals convenience function for
// string comparison BoolFn
func StringEquals(compareTo string) BoolFn[string] {
	return func(s string, _ int, _ []string) bool {
		return compareTo == s
	}
}

// Pipeline func composes a transform pipeline around a string
func Pipeline[T any](fns ...func(T) T) func(T) T {
	return func(item T) T {
		out := item

		for _, fn := range fns {
			out = fn(out)
		}

		return out
	}
}
