package stringslice

func Some(strs []string, someFn func(string) bool) bool {
	for _, str := range strs {
		if someFn(str) {
			return true
		}
	}

	return false
}

func Map(strs []string, mapFn func(string) string) []string {
	mapped := make([]string, len(strs), len(strs))

	for i := 0; i < len(strs); i++ {
		mapped[i] = mapFn(strs[i])
	}

	return mapped
}

func MapError(strs []string, mapErrFn func(string) (string, error)) []string {
	mapped := make([]string, 0, len(strs))

	for _, str := range strs {
		result, err := mapErrFn(str)
		if err != nil {
			continue
		}

		mapped = append(mapped, result)
	}

	return mapped
}
