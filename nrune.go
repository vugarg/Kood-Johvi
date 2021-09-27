package piscine

func NRune(s string, n int) rune {
	runedS := []rune(s)
	len := len(runedS)

	if n > len || n <= 0 {
		return 0
	} else {
		return runedS[n-1]
	}
}
