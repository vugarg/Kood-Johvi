package piscine

func AlphaCount(s string) int {
	runedS := []rune(s)
	counter := 0
	for _, word := range runedS {
		if (word > 64 && word < 91) || (word > 96 && word < 123) {
			counter++
		}
	}
	return counter
}
