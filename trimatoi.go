package piscine

func TrimAtoi(s string) int {
	result := 0
	sign := 1
	foundDigit := false

	for i := 0; i < len(s); i++ {
		if s[i] >= '0' && s[i] <= '9' {
			result = result*10 + int(s[i]-'0')
			foundDigit = true
		} else if s[i] == '-' && !foundDigit {
			sign = -1
		} else if s[i] == '+' && !foundDigit {
			sign = 1
		}
	}

	if !foundDigit {
		return 0
	}

	return result * sign
}
