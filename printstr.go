package piscine

import (
	"github.com/01-edu/z01"
)

func PrintStr(s string) {
	for _, m := range s {
		if m == '\n' {
			z01.PrintRune('\\')
			z01.PrintRune('n')
		} else {
			z01.PrintRune(m)
		}
	}
}
