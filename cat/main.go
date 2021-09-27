package main

import (
	"io/ioutil"
	"os"

	"github.com/01-edu/z01"
)

func printString(inp string) {
	for _, v := range inp {
		z01.PrintRune(v)
	}
}

func printFile(filename string) {
	dat, err := ioutil.ReadFile(filename)
	if err != nil {
		printString("ERROR: open " + filename + ": no such file or directory\n")
		os.Exit(1)
	} else {
		for _, v := range dat {
			z01.PrintRune(rune(v))
		}
	}
}

func printStdin() {
	scanner, _ := ioutil.ReadAll(os.Stdin)
	for _, v := range scanner {
		z01.PrintRune(rune(v))
	}
}

func main() {
	if len(os.Args) > 1 {
		for _, f := range os.Args[1:] {
			printFile(f)
		}
	} else {
		printStdin()
	}
}
