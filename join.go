package piscine

func Join(strs []string, sep string) string {
	if len(strs) == 0 {
		return ""
	}

	result := strs[0]
	for y := 1; y < len(strs); y++ {
		result += sep + strs[y]
	}
	return result
}
