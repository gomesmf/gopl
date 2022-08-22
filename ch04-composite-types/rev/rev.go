package main

import "fmt"

// reverse reverses a slice of ints in place
func reverse(s []int) {
	for i, j := 0, len(s)-1; i < j; i, j = i+1, j-1 {
		s[i], s[j] = s[j], s[i]
	}
}

func main() {
	a := [...]int{0, 1, 2, 3, 4, 5}
	reverse(a[:])
	fmt.Println(a)

	b := []int{0, 1, 2, 3, 4, 5}
	fmt.Printf("b: %v\n", b)
	//Rotate s left by two positions.
	reverse(b[:2])
	fmt.Printf("b: %v\n", b)
	reverse(b[2:])
	fmt.Printf("b: %v\n", b)
	reverse(b)
	fmt.Printf("b: %v\n", b)

	fmt.Printf("cap(b): %d; len(b): %d\n", cap(b), len(b))

	var c []int
	fmt.Println(c, len(c), c == nil)
}
