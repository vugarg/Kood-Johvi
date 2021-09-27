package main

import (
	"github.com/01-edu/z01"
)

type point struct {
	x, y int
}

func setPoint(ptr *point) {
	ptr.x = 42
	ptr.y = 21
}

func runeToInt(number rune) int {
	count := 0
	for i := '1'; i <= number; i++ {
		count++
	}
	return count
}

func main() {
	points := &point{}

	setPoint(points)

	var result string = "x = 42, y = 21"

	r := []rune(result)
	printer(r)
}

func printer(result []rune) {
	for _, ch := range result {
		z01.PrintRune(ch)
	}
	z01.PrintRune('\n')
}
