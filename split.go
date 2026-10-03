package piscine

func Split(s, sep string) []string {
	var result []string
	sepLen := len(sep)

	if sepLen == 0 {
		for _, c := range s {
			result = append(result, string(c))
		}
		return result
	}

	start := 0
	for i := 0; i <= len(s)-sepLen; i++ {
		if s[i:i+sepLen] == sep {
			result = append(result, s[start:i])
			start = i + sepLen
			i += sepLen - 1
		}
	}
	result = append(result, s[start:])
	return result
}
