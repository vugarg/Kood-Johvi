package piscine

func SortWordArr(a []string) {
	var outOfarr string

	i := 1
	for i < len(a) {
		if a[i-1] > a[i] {
			outOfarr = a[i]
			a[i] = a[i-1]
			a[i-1] = outOfarr
			i = 1
		} else {
			i++
		}
	}
}
