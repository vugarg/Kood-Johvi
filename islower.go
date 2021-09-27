package piscine

func IsLower(s string) bool {
	for _, ch := range []rune(s) {
		if !(ch >= 'a' && ch <= 'z') {
			return false
		}
	}
	return true
}
