package main

import (
	"os"

	"github.com/01-edu/z01"
)

func main() {
	name := os.Args
	len := len(name)

	for i := len - 1; i > 0; i-- {
		for _, i := range name[i] {
			z01.PrintRune(i)
		}
		z01.PrintRune(10)
	}
}
