package piscine

func Fibonacci(index int) int {
	res := 0
	if index < 0 {
		return -1
	} else if index < 2 {
		return index
	} else {
		res = (Fibonacci(index-1) + Fibonacci(index-2))
	}
	return res
}

// Write a recursive function that returns the value of the fibonacci sequence matching the index passed as parameter.

// The first value is at index 0.

// The sequence starts this way: 0, 1, 1, 2, 3 etc...

// A negative index will return -1.

// for is forbidden for this exercise.
