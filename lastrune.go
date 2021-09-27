package piscine

func LastRune(s string) rune {
	runedS := []rune(s)
	len := len(runedS)
	return runedS[len-1]
}
