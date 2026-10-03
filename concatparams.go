package piscine

func ConcatParams(args []string) string {
	result := ""
	for i, s := range args {
		result += s
		if i < len(args)-1 {
			result += "\n"
		}
	}
	return result
}
