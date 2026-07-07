package piscine

func BasicAtoi(s string) int {
	var k int
	for _, m := range s {
		k = k*10 + int(m-'0')
	}
	return k
}
