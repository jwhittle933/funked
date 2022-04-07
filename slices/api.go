package slices

import "github.com/jwhittle933/gonads/option"

// SlicedAssertion convenience interface for
// all slice funcs that return a bool
type SlicedAssertion[T any] interface {
	Some(Slice[T], BoolFn[T]) bool
	Every(Slice[T], BoolFn[T]) bool
	Empty(Slice[T]) bool
}

// SlicedItem convenience interface for
// all lice funcs that return an item from the Slice
type SlicedItem[T any] interface {
	Find(Slice[T], BoolFn[T]) option.Option[T]
	FindIndex(Slice[T], BoolFn[T]) option.Option[int]
	First(Slice[T]) option.Option[T]
	Last(Slice[T]) option.Option[T]
	At(Slice[T], int) option.Option[T]
}

// Slicer convenience interface for all
// slice funcs that return another Slice
type Slicer[T any] interface {
	Filter(Slice[T], BoolFn[T]) Slice[T]
	Map(Slice[T], TFn[T]) Slice[T]
	Prepend(Slice[T], T) Slice[T]
}

// Sliced convenience interface for Bare API
type Sliced[T any] interface {
	Slicer[T]
	SlicedAssertion[T]
	SlicedItem[T]
}

func Filter[T any](s Slice[T], bfn BoolFn[T]) Slice[T] {
	return s.Filter(bfn)
}

func Map[T any](s Slice[T], sfn TFn[T]) Slice[T] {
	return s.Map(sfn)
}

func Some[T any](s Slice[T], bfn BoolFn[T]) bool {
	return s.Some(bfn)
}

func Every[T any](s Slice[T], bfn BoolFn[T]) bool {
	return s.Every(bfn)
}

func Find[T any](s Slice[T], bfn BoolFn[T]) option.Option[T] {
	return s.Find(bfn)
}

func FindIndex[T any](s Slice[T], bfn BoolFn[T]) option.Option[int] {
	return s.FindIndex(bfn)
}

func Last[T any](s Slice[T]) option.Option[T] {
	return s.Last()
}

func First[T any](s Slice[T]) option.Option[T] {
	return s.First()
}

func At[T any](s Slice[T], index int) option.Option[T] {
	return s.At(index)
}

func Prepend[T any](s Slice[T], item T) Slice[T] {
	return s.Prepend(item)
}

func Empty[T any](s Slice[T]) bool {
	return s.Empty()
}
