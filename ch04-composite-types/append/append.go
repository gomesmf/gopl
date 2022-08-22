package main

import "fmt"

func appendInt(x []int, y int) []int {
	var z []int
	zlen := len(x) + 1
	if zlen <= cap(x) {
		// There is room to grow. Extend the slice.
		z = x[:zlen]
	} else {
		// There is insufficient space. Allocate a new array.
		// Grow by doubling, for amortized linear complexity.
		zcap := zlen
		if zcap < 2*len(x) {
			zcap = 2 * len(x)
		}
		z = make([]int, zlen, zcap)
		copy(z, x) // a built-in function
	}
	z[len(x)] = y
	return z
}

func main() {
	var x []int
	fmt.Println(x, len(x), cap(x))
	x = appendInt(x, 1)
	fmt.Println(x, len(x), cap(x))
	x = appendInt(x, 2)
	fmt.Println(x, len(x), cap(x))
	x = appendInt(x, 3)
	fmt.Println(x, len(x), cap(x))
	x = appendInt(x, 4)
	fmt.Println(x, len(x), cap(x))
	x = appendInt(x, 5)
	fmt.Println(x, len(x), cap(x))
}
