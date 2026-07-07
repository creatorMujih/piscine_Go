package piscine

func IsPrintable(str string) bool {
	for _, p := range str {
		if p >= ' ' && p <= '~' {
			continue
		} else {
			return false
		}
	}
	return true
}
