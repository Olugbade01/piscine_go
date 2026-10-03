package piscine

func Capitalize(s string) string {
	runes := []rune(s)
	newWord := true

	for i := 0; i < len(runes); i++ {
		c := runes[i]

		// check if alphanumeric
		isAlphaNum := (c >= 'a' && c <= 'z') ||
			(c >= 'A' && c <= 'Z') ||
			(c >= '0' && c <= '9')

		if isAlphaNum {
			if newWord {
				// to upper if lowercase letter
				if c >= 'a' && c <= 'z' {
					runes[i] = c - 32
				}
				newWord = false
			} else {
				// to lower if uppercase letter
				if c >= 'A' && c <= 'Z' {
					runes[i] = c + 32
				}
			}
		} else {
			// punctuation resets word
			newWord = true
		}
	}

	return string(runes)
}
