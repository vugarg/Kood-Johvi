package piscine

func IsPrintable(s string) bool {
	for _, ch := range []rune(s) {
		if ch >= 0 && ch <= 31 {
			return false
		}
	}
	return true
}
