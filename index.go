package piscine

func Index(s string, toFind string) int {
	runeS := []rune(s)
	runeToFind := []rune(toFind)
	n := len(runeToFind)

	if n == 0 {
		return 0
	} else if n == 1 {
		for i, ch := range runeS {
			if ch == runeToFind[0] {
				return i
			}
		}
	}

	for i := range runeS {
		if runeToFind[0] == runeS[i] {
			tmp := false
			for j := range runeToFind {
				if runeS[i+j] == runeToFind[j] {
					tmp = true
				} else {
					tmp = false
				}
			}
			if tmp {
				return i
			}
		}
	}
	return -1
}
