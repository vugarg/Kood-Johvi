package piscine

func Atoi(s string) int {
	var zero rune = 48
	runes := []rune(s)
	result := 0
	length := len([]rune(s))

	if length == 0 {
		return 0
	}

	var check bool

	if runes[0] == 45 {
		runes[0] = 48
		check = true
	}

	for _, word := range runes {
		if word < 48 || word > 57 {
			return 0
		} else {
			a := int(word - zero)
			result = result*10 + a
		}
	}

	if check {
		result *= -1
	}
	return result
}
