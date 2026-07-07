package piscine

func CountIf(f func(string) bool, a []string) int {
	count := 0
	for _, x := range a {
		if f(x) == true {
			count++
		}
	}
	return count
}
