package intslice

func Filter(ints []int, bfn BoolFn) []int {
	return From(ints).Filter(bfn)
}

func Map(ints []int, ifn IntFn) []int {
	return From(ints).Map(ifn)
}

func Some(ints []int, bfn BoolFn) bool {
	return From(ints).Some(bfn)
}

func Every(ints []int, bfn BoolFn) bool {
	return From(ints).Every(bfn)
}

func Find(ints []int, bfn BoolFn) *int {
	return From(ints).Find(bfn)
}

func FindIndex(ints []int, bfn BoolFn) *int {
	return From(ints).FindIndex(bfn)
}

func Includes(ints []int, includes int) bool {
	return From(ints).Includes(includes)
}

func IndexOf(ints []int, integer int) *int {
	return From(ints).IndexOf(integer)
}

func Join(ints []int, sep string) string {
	return From(ints).Join(sep)
}

func Last(ints []int) *int {
	return From(ints).Last()
}

func First(ints []int) *int {
	return From(ints).First()
}

func Prepend(ints []int, i int) []int {
	return From(ints).Prepend(i)
}

func At(ints []int, index int) *int {
	return From(ints).At(index)
}

func Empty(ints []int) bool {
	return From(ints).Empty()
}

func Sort(ints []int) []int {
	return From(ints).Sort()
}


