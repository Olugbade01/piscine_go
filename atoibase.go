package piscine

func AtoiBase(s string, base string) int {
	// Validate base
	if len(base) < 2 {
		return 0
	}
	for i := 0; i < len(base); i++ {
		if base[i] == '+' || base[i] == '-' {
			return 0
		}
		for j := i + 1; j < len(base); j++ {
			if base[i] == base[j] {
				return 0
			}
		}
	}

	baseLen := len(base)
	result := 0

	for i := 0; i < len(s); i++ {
		// Find the value of the current character in the base
		digitVal := -1
		for j := 0; j < baseLen; j++ {
			if s[i] == base[j] {
				digitVal = j
				break
			}
		}
		if digitVal == -1 {
			return 0
		}
		result = result*baseLen + digitVal
	}

	return result
}
