package piscine

func StringToIntSlice(str string) []int {
	var result []int
	for _, i := range str {
		result = append(result, int(i))
	}
	return result
}
