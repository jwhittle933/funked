package stringslice

type StringFunked interface {
	Filter([]string) []string
	Map([]string) []string
	Some([]string) bool
}

func Filter(strs []string, bfn BoolFn) []string {
	return From(strs).Filter(bfn)
}

func Map(strs []string, sfn StringFn) []string {
	return From(strs).Map(sfn)
}

func Some(strs []string, bfn BoolFn) bool {
	return From(strs).Some(bfn)
}

func Every(strs []string, bfn BoolFn) bool {
	return From(strs).Every(bfn)
}
