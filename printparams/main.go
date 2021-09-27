package main

import (
	"os"

	"github.com/01-edu/z01"
)

func main() {
	name := os.Args
	len := len(name)

	for i := 1; i < len; i++ {
		for _, i := range name[i] {
			z01.PrintRune(i)
		}
		z01.PrintRune(10)
	}
}
