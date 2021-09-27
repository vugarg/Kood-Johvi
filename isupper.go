package piscine

func IsUpper(s string) bool {
	for _, ch := range []rune(s) {
		if !(ch >= 65 && ch <= 90) {
			return false
		}
	}
	return true
}
