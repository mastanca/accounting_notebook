package misc

func TransposeAndFlatten(input [][]string) []string {
	var level int
	for _, innerSlice := range input {
		currentLen := len(innerSlice)
		if currentLen > level {
			level = currentLen
		}
	}

	var result []string
	for i := 0; i < level; i++ {
		for _, innerSlice := range input {
			if i <= len(innerSlice)-1 {
				result = append(result, innerSlice[i])
			}
		}
	}
	return result
}
