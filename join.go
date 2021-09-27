package piscine

func Join(strs []string, sep string) string {
	result := ""
	len := len(strs)
	for _, word := range strs {
		if word == strs[len-1] {
			result += word
		} else {
			result = result + word + sep
		}
	}
	return result
}
