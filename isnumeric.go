package piscine

func IsNumeric(str string) bool {
	for _, ch := range str {
		if !(ch >= '0' && ch <= '9') {
			return false
		}
	}
	return true
}
