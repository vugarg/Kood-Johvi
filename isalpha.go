package piscine

func IsAlpha(s string) bool {
	for _, ch := range []rune(s) {
		if !((ch >= 'a' && ch <= 'z') || (ch >= 'A' && ch <= 'Z') || (ch >= '0' && ch <= '9')) {
			return false
		}
	}
	return true
}
