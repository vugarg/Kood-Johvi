package piscine

func BasicAtoi2(s string) int {
	var zero rune = 48
	runes := []rune(s)
	result := 0

	for _, word := range runes {
		if word < 48 || word > 57 {
			return 0
		} else {
			a := int(word - zero)
			result = result*10 + a
		}
	}
	if runes[0] == '-' {
		result *= -1
	}
	return result
}
