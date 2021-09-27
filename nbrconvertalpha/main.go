package main

import (
	"os"

	"github.com/01-edu/z01"
)

func main() {
	args := os.Args[1:]
	var abc []rune
	var ABC []rune
	b := false

	for i := 'a'; i <= 'z'; i++ {
		abc = append(abc, i)
	}

	for i := 'A'; i <= 'Z'; i++ {
		ABC = append(ABC, i)
	}

	for _, e := range args {

		if e == "--upper" {
			b = true
			continue
		}

		if Atoi(e) > 26 || Atoi(e) < 1 {
			z01.PrintRune(' ')
			continue
		}

		if b {
			z01.PrintRune(rune(ABC[Atoi(e)-1]))
		} else {
			z01.PrintRune(rune(abc[Atoi(e)-1]))
		}
	}
	if len(args) > 1 {
		z01.PrintRune('\n')
	}
}

func Atoi(s string) int {
	var result int
	var check bool
	r := []rune(s)

	if len(r) == 0 {
		return 0
	}

	if r[0] == 45 {
		r[0] = 48
		check = true
	} else if r[0] == 43 {
		r[0] = 48
	}

	for _, e := range r {
		if e < 48 || e > 57 {
			return 0
		}
		a := int(e - 48)
		result = result*10 + a
	}

	if check {
		result *= -1
	}

	return result
}
