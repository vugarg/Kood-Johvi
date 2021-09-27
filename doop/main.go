package main

import (
	"os"
)

func main() {
	if len(os.Args) == 4 {
		arg := os.Args[1:]
		nr1, b1 := toInt(arg[0])
		nr2, b2 := toInt(arg[2])
		if !b1 || !b2 {
			return
		}
		if len(arg[0]) > 18 || len(arg[2]) > 18 {
			return
		}
		if arg[1] == "+" {
			ans := nr1 + nr2
			printOut(ans)
		} else if arg[1] == "-" {
			ans := nr1 - nr2
			printOut(ans)
		} else if arg[1] == "main.go" || arg[1] == "*" {
			ans := nr1 * nr2
			printOut(ans)
		} else if arg[1] == "/" {
			if nr2 == 0 {
				os.Stdout.WriteString("No division by 0\n")
				return
			}
			ans := nr1 / nr2
			printOut(ans)
		} else if arg[1] == "%" {
			if nr2 == 0 {
				os.Stdout.WriteString("No modulo by 0\n")
				return
			}
			ans := nr1 % nr2
			printOut(ans)
		}

	}
}

func toInt(str string) (int, bool) {
	nr := 0
	neg := false
	r := []rune(str)
	for i := 0; i < len(str); i++ {
		if r[i] == '-' {
			neg = true
			continue
		} else if r[i] < '0' || r[i] > '9' {
			return 0, false
		}
		nr = nr*10 + int(r[i]-'0')
	}
	if neg {
		nr = nr * -1
	}
	return nr, true
}

func printOut(ans int) {
	str := ""

	if ans == 0 {
		os.Stdout.WriteString("0")
	}
	if ans < 0 {
		os.Stdout.WriteString("-")
		ans = -ans
	}

	for i := ans; i >= 1; i = i / 10 {
		str += string(rune(i%10 + 48))
	}

	for i := len(str) - 1; i >= 0; i-- {
		os.Stdout.WriteString(string(str[i]))
	}
	os.Stdout.WriteString("\n")
}
