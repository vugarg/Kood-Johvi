package main

import (
	"os"

	"github.com/01-edu/z01"
)

func main() {
	name := os.Args
	for _, i := range name[0] {
		if !(i == '.' || i == '/') {
			z01.PrintRune(i)
		}
	}
	z01.PrintRune(10)
}
