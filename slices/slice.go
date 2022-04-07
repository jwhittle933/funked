package slices

import "github.com/jwhittle933/gonads/option"

type Slice[T any] []T

func From[T any](strs ...T) Slice[T] {
	return strs
}

func (s Slice[T]) Filter(bfn BoolFn[T]) Slice[T] {
	return bfn.Filter(s)
}

func (s Slice[T]) Map(sfn TFn[T]) Slice[T] {
	return sfn.Map(s)
}

func (s Slice[T]) Some(bfn BoolFn[T]) bool {
	return bfn.Some(s)
}

func (s Slice[T]) Every(bfn BoolFn[T]) bool {
	return bfn.Every(s)
}

func (s Slice[T]) Find(bfn BoolFn[T]) option.Option[T] {
	return bfn.Find(s)
}

func (s Slice[T]) FindIndex(bfn BoolFn[T]) option.Option[int] {
	return bfn.FindIndex(s)
}

func (s Slice[T]) Last() option.Option[T] {
	if s.Empty() {
		return option.None[T]()
	}

	return option.Some(s[len(s)-1])
}

func (s Slice[T]) First() option.Option[T] {
	if s.Empty() {
		return option.None[T]()
	}

	return option.Some(s[0])
}

func (s Slice[T]) At(index int) option.Option[T] {
	if index < 0 || len(s)-1 < index {
		return option.None[T]()
	}

	return option.Some(s[index])
}

func (s Slice[T]) Prepend(i T) Slice[T] {
	out := make([]T, 1, len(s)+1)
	out[0] = i
	return append(out, s...)
}

func (s Slice[T]) Empty() bool {
	return len(s) == 0
}

func (s Slice[T]) ToMap() map[int]T {
	out := make(map[int]T, len(s))

	for i, str := range s {
		out[i] = str
	}

	return out
}
