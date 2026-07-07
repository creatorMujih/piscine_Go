package piscine

func IsAlpha(s string) bool {
	if len(s) == 0 {
		return true
	}
	for _, k := range s {
		if (k >= 'a' && k <= 'z') || (k >= 'A' && k <= 'Z') || (k >= '0' && k <= '9') {
			continue
		} else {
			return false
		}
	}
	return true
}
