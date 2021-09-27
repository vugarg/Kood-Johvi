package main

import (
	"os"

	"github.com/01-edu/z01"
)

func printStr(s string) {
	for _, r := range s {
		z01.PrintRune(r)
	}
	z01.PrintRune('\n')
}

func isEven(nbr int) bool {
	if nbr%2 == 1 {
		return false
	}
	return true
}

func main() {
	lengthOfArg := getLen(os.Args[1:])
	EvenMsg := "I have an even number of arguments"
	OddMsg := "I have an odd number of arguments"

	if isEven(lengthOfArg) {
		printStr(EvenMsg)
	} else {
		printStr(OddMsg)
	}
}

func getLen(args []string) int {
	count := 0
	for range args {
		count++
	}
	return count
}

func getStrLen(str []rune) int {
	count := 0
	for range str {
		count++
	}
	return count
}
