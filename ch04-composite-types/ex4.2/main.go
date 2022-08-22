package main

import (
	"crypto/sha256"
	"crypto/sha512"
	"flag"
	"fmt"
)

func main() {
	hash512 := flag.Bool("sha512", false, "prints the SHA512 of its stdin")
	hash384 := flag.Bool("sha384", false, "prints the SHA384 of its stdin")
	flag.Parse()

	var input []byte

	fmt.Scan(&input)

	switch {
	case *hash512:
		sum := sha512.Sum512(input)
		fmt.Printf("sha512:%x\n", sum)
	case *hash384:
		sum := sha512.Sum384(input)
		fmt.Printf("sha384:%x\n", sum)
	default:
		sum := sha256.Sum256(input)
		fmt.Printf("sha256:%x\n", sum)
	}
}
