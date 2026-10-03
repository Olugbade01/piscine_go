package piscine

// Map applies f to each element of a and returns a slice with the results.
func Map(f func(int) bool, a []int) []bool {
	if a == nil {
		return nil
	}
	res := make([]bool, len(a))
	for i, v := range a {
		res[i] = f(v)
	}
	return res
}
