package piscine

func IsNumeric(str string) bool {
	for _, m := range str {
		if m >= '0' && m <= '9' {
			continue
		} else {
			return false
		}
	}
	return true
}
