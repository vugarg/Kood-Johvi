package piscine

func TrimAtoi(s string) int {
	var zero rune = 48
	runes := []rune(s)
	nm := 0

	var minus bool
	for _, ch := range s {
		if ch >= '0' && ch <= '9' {
			break
		}

		if ch == '-' {
			minus = true
		}
	}

	for _, ch := range runes {
		if ch >= '0' && ch <= '9' {
			a := int(ch - zero)
			nm = nm*10 + a
		}
	}

	if minus {
		nm *= -1
	}
	return nm
}
