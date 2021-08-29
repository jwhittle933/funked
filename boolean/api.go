package boolean

type BoolComposer func(first, second bool) bool

func Compose(composers ...BoolComposer) BoolComposer {
	return func(first, second bool) bool {
		var out bool

		for _, c := range composers {
			out = c(first, second)
		}

		return out
	}
}

func And(first, second bool) bool {
	return first && second
}

func AndNot(first, second bool) bool {
	return first && !second
}

func Or(first, second bool) bool {
	return first || second
}

func OrNot(first, second bool) bool {
	return first || !second
}
