package piscine

func UltimateDivMod(a *int, b *int) {
	var a1 int = *a
	var b1 int = *b
	*a = a1 / b1
	*b = a1 % b1
}

// This function will divide the int a and b.
// The result of this division will be stored in the int pointed by a.
// The remainder of this division will be stored in the int pointed by b.
