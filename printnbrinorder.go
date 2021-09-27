package piscine

import "github.com/01-edu/z01"

func PrintNbrInOrder(n int) {
	var digits []rune
	result := 0

	if n < 0 || n == 0 {
		z01.PrintRune('0')
	}

	for n > 0 {
		digit := n % 10
		n /= 10

		digits = append(digits, rune(digit))
	}

	for i := 0; i < len(digits)-1; i++ {
		for j := 0; j < len(digits)-i-1; j++ {
			if digits[j] > digits[j+1] {
				digits[j], digits[j+1] = digits[j+1], digits[j]
			}
		}
	}

	for _, word := range digits {
		result = result*10 + int(word)
	}

	for i := '0'; i <= '9'; i++ {
		for _, digit := range digits {
			if i == digit+48 {
				z01.PrintRune(i)
			}
		}
	}
}
