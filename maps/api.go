package maps

func StringKeys(m map[string]interface{}) []string {
	out := make([]string, 0)

	for k := range m {
		out = append(out, k)
	}

	return out
}

func IntKeys(m map[int]interface{}) []int {
	out := make([]int, 0)

	for k := range m {
		out = append(out, k)
	}

	return out
}
