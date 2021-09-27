package piscine

import (
	"github.com/01-edu/z01"
)

func PrintStr(s string) {
	for _, word := range s {
		z01.PrintRune(word)
	}
	z01.PrintRune('\n')
}

// Write a function that prints one by one the characters of a string on the screen.
