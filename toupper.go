package piscine

func ToUpper(s string) string {
	runes := []rune(s)
	for i, ch := range runes {
		if ch >= 'a' && ch <= 'z' {
			runes[i] -= 32
		}
	}
	return string(runes)
}
