package main

import (
	"fmt"
	"math"

	"github.com/jwhittle933/funked/slices/intslice"
)

func main() {
	ints := []int{10, 1005, 24, 30, 44, 503}

	fmt.Printf("\033[0;34mintslice.From(x).Map(y) ->\033[0m %+v\n", intslice.From(ints).Map(inc))
	fmt.Printf("\033[0;34mintslice.IntFn(y).Map(x) ->\033[0m %+v\n\n", intslice.IntFn(inc).Map(ints))
	fmt.Printf("\033[0;34mintslice.Map(x, y) ->\033[0m %+v\n", intslice.Map(ints, inc))
	fmt.Printf("\033[0;34mintslice.Includes(x, y) ->\033[0m %+v\n", intslice.Includes(ints, 1005))
	fmt.Printf("\033[0;34mintslice.IndexOf(x, y) ->\033[0m %d\n", *intslice.IndexOf(ints, 1005))
	fmt.Printf("\033[0;34mintslice.Join(x, y) ->\033[0m %+v\n", intslice.Join(ints, ", "))
	fmt.Printf("\033[0;34mintslice.First(x) ->\033[0m %d\n", *intslice.First(ints))
	fmt.Printf("\033[0;34mintslice.Last(x) ->\033[0m %d\n", *intslice.Last(ints))
	fmt.Printf("\033[0;34mintslice.Empty(x) ->\033[0m %+v\n", intslice.Empty(ints))
	fmt.Printf("\033[0;34mintslice.Prepend(x, y) ->\033[0m %+v\n", intslice.Prepend(ints, 500))
	fmt.Printf("\033[0;34mintslice.Sort(x) ->\033[0m %+v\n\n", intslice.Sort(ints))

	mapped := intslice.
		IntFn(inc).
		And(times(10)).
		And(dec).
		Map(ints)

	fmt.Printf("\033[0;34mintslice.IntFn(inc).And(times(10)).And(dec).Map(ints) ->\033[0m %+v\n", mapped)

	filtered := intslice.
		BoolFn(mod10).
		AndNot(gt(4)).
		And(prime).
		Filter(mapped)

	fmt.Printf("\033[0;34mintslice.BoolFn(mod10).AndNot(gt(4)).And(prime).Filter(mapped) ->\033[0m %+v\n", filtered)
}

func inc(i int, _ int, _ []int) int {
	return i + 1
}

func times(delta int) intslice.IntFn {
	return func(i int, _ int, _ []int) int {
		return i * delta
	}
}

func dec(i int, _ int, _ []int) int {
	return i - 1
}

func mod10(i int, _ int, _ []int) bool {
	return i%10 != 0
}

func gt(delta int) intslice.BoolFn {
	return func(_ int, iter int, _ []int) bool {
		return iter > delta
	}
}

func prime(val int, _ int, _ []int) bool {
	for i := 2; i <= int(math.Floor(float64(val)/2)); i++ {
		if val%i == 0 {
			return false
		}
	}

	return val > 1
}
