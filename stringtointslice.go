package piscine

func StringToIntSlice(str string) []int {
	if str == "" {
		return nil
	}

	result := make([]int, 0)

	for _, r := range str {
		result = append(result, int(r))
	}

	return result
}
