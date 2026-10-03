package piscine

func ShoppingSummaryCounter(str string) map[string]int {
	result := make(map[string]int)
	item := ""

	for _, word := range str {
		if word == ' ' {
			result[item]++
			item = ""
		} else {
			item += string(word)
		}
	}
	result[item]++
	return result
}
