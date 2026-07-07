package piscine

func Compact(ptr *[]string) int {
	var newList []string
	for _, y := range *ptr {
		if y != "" {
			newList = append(newList, y)
		}
	}
	*ptr = newList
	return len(newList)
}
