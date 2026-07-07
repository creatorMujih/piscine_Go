package piscine

func ToLower(s string) string {
	var result string
	for _, f := range s {
		if f >= 'A' && f <= 'Z' {
			result += string(f + 32)
		} else {
			result += string(f)
		}
	}
	return result
}
