package piscine

func IterativeFactorial(nb int) int {
	res := 1
	if nb < 0 || nb > 20 {
		return 0
	} else if nb == 0 {
		return 1
	} else {
		for i := 1; i <= nb; i++ {
			res = i * res
		}
	}
	return res
}
