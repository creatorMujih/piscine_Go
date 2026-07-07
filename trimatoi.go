package piscine

func TrimAtoi(s string) int {
	sign := 1
	started := false
	result := 0

	for _, d := range s {
		if d == '-' && !started {
			sign = -1
		}
		if d >= '0' && d <= '9' {
			started = true
			result = result*10 + int(d-'0')
		}
	}

	return result * sign
}
