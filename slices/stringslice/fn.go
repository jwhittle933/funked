package stringslice

type BoolFn func(string) bool
type StringFn func(string) string

func (bfn BoolFn) Filter(strs []string) []string {
	filtered := make([]string, 0, len(strs))

	for _, str := range strs {
		if bfn(str) {
			filtered = append(filtered, str)
		}
	}

	return filtered
}

func (sfn StringFn) Map(strs []string) []string {
	length := len(strs)
	mapped := make([]string, length, length)

	for i := 0; i < length; i++ {
		mapped[i] = sfn(strs[i])
	}

	return mapped
}

func (bfn BoolFn) Some(strs []string) bool {
	for _, str := range strs {
		if bfn(str) {
			return true
		}
	}

	return false
}

func (bfn BoolFn) Every(strs []string) bool {
	for _, str := range strs {
		if !bfn(str) {
			return false
		}
	}

	return true
}

func (sfn StringFn) With(next StringFn) StringFn {
	return func(s string) string {
		return next(sfn(s))
	}
}

func (bfn BoolFn) With(next BoolFn) BoolFn {
	return func(s string) bool {
		if bfn(s) {
			return next(s)
		}

		return false
	}
}
