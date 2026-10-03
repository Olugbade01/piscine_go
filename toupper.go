package piscine

func ToUpper(s string) string {
	result := []rune(s)

	for i, r := range result {
		if r >= 'a' && r <= 'z' {
			result[i] = r - 32
		}
	}

	return string(result)
}
