package piscine

import "github.com/01-edu/z01"

func PrintComb() {
	const Zero = 48
	for i := 0; i <= 9; i++ {
		for j := i + 1; j <= 9; j++ {
			for k := j + 1; k <= 9; k++ {
				z01.PrintRune(rune(Zero + i))
				z01.PrintRune(rune(Zero + j))
				z01.PrintRune(rune(Zero + k))
				if !(i == 7 && j == 8 && k == 9) {
					z01.PrintRune(',')
					z01.PrintRune(' ')
				}
			}
		}
	}
	z01.PrintRune('\n')
}
