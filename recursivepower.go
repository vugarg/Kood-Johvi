package piscine

func RecursivePower(nb int, power int) int {
	res := 1
	if power < 0 {
		return 0
	}
	if nb == 0 && power == 0 {
		return 1
	} else if nb == 0 {
		return 0
	} else if power >= 1 {
		res = nb * res * RecursivePower(nb, power-1)
	}
	return res
}
