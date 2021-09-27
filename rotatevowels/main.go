package main

import (
	"os"

	"github.com/01-edu/z01"
)

func main() {
	args := ArgsMerge(os.Args)

	r := []rune(args)
	if len(args) > 0 {

		vowels := GetVowels(r)
		vowels = reverseVowels(vowels)

		j := 0
		for i := 0; i < len(r); i++ {
			if Vowel(r[i]) {
				r[i] = vowels[j]
				j++
			}
		}

		for i := 0; i < len(r); i++ {
			z01.PrintRune(r[i])
		}
		z01.PrintRune('\n')
	} else {
		z01.PrintRune('\n')
	}
}

func reverseVowels(r []rune) []rune {
	run := []rune{}
	for i := len(r) - 1; i >= 0; i-- {
		run = append(run, r[i])
	}
	return run
}

func ArgsMerge(s []string) string {
	str := ""
	for i := 1; i < len(s); i++ {
		str += s[i]
		str += " "
	}
	return str
}

func Vowel(r rune) bool {
	if r == 'a' || r == 'e' || r == 'i' || r == 'o' ||
		r == 'u' || r == 'A' || r == 'E' || r == 'I' || r == 'O' ||
		r == 'U' {
		return true
	}
	return false
}

func GetVowels(r []rune) []rune {
	vowels := []rune{}
	for i := 0; i < len(r); i++ { // GET VOWELS
		if Vowel(r[i]) {
			vowels = append(vowels, r[i])
		}
	}
	return vowels
}
