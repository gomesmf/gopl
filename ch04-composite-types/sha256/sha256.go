package main

import (
	"crypto/sha256"
	"fmt"
)

func main() {
	c1 := sha256.Sum256([]byte("x"))
	c2 := sha256.Sum256([]byte("X"))
	fmt.Printf("%T %[1]x %[1]v\n", c1)
	fmt.Printf("%T %[1]x %[1]v\n", c2)

}
