package piscine

func ActiveBits(a int) int {
	count := 0
	for a != 0 {
		if a%2 == 1 {
			count++
		}
		a /= 2
	}
	return count
}
