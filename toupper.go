package piscine

func ToUpper(s string) string {
	var result string
	for _, y := range s {
		if y >= 'a' && y <= 'z' {
			result += string(y - 32)
		} else {
			result += string(y)
		}
	}
	return result
}
