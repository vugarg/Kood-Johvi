package piscine

func ToLower(s string) string {
	runes := []rune(s)
	for i, ch := range runes {
		if ch >= 'A' && ch <= 'Z' {
			runes[i] += 32
		}
	}
	return string(runes)
}
