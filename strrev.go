package piscine

func StrRev(s string) string {
	result := ""
	for _, v := range s {
		result = string(v) + result
	}
	return result
}

// Write a function that reverses a string.

// This function will return the reversed string.
