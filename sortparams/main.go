package main

import (
	"os"

	"github.com/01-edu/z01"
)

func main() {
	name := os.Args
	len := len(name)

	for i := 1; i <= len-1; i++ {
		for j := 1; j <= len-i-1; j++ {
			if name[j] > name[j+1] {
				name[j], name[j+1] = name[j+1], name[j]
			}
		}
	}

	for i := 1; i < len; i++ {
		for _, i := range name[i] {
			z01.PrintRune(i)
		}
		z01.PrintRune(10)
	}
}
