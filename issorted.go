package piscine

func IsSorted(f func(a1, a2 int) int, tab []int) bool {
	var sorted bool = true
	for x := 0; x < len(tab)-1; x++ {
		if f(tab[x], tab[x+1]) > 0 {
			sorted = false
		}
	}
	if sorted == false {
		sorted = true
		for x := 0; x < len(tab)-1; x++ {
			if f(tab[x], tab[x+1]) < 0 {
				sorted = false
			}
		}
	}
	return sorted
}
