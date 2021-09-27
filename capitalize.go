package piscine

func Capitalize(s string) string {
	runes := []rune(s)
	length := len(runes)

	if runes[0] >= 'a' && runes[0] <= 'z' {
		runes[0] -= 32
	}

	for i := 1; i < length; i++ {
		if runes[i] >= 'A' && runes[i] <= 'Z' {
			if runes[i-1] >= 'A' && runes[i-1] <= 'Z' ||
				runes[i-1] >= 'a' && runes[i-1] <= 'z' ||
				runes[i-1] >= '0' && runes[i-1] <= '9' {
				runes[i] += 32
			}
		} else if runes[i] >= 'a' && runes[i] <= 'z' {
			if runes[i-1] >= 'A' && runes[i-1] <= 'Z' ||
				runes[i-1] >= 'a' && runes[i-1] <= 'z' ||
				runes[i-1] >= '0' && runes[i-1] <= '9' {
				continue
			} else {
				runes[i] -= 32
			}
		}
	}

	return string(runes)
}
