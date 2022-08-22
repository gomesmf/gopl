package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	wcount := make(map[string]int)
	input := bufio.NewScanner(os.Stdin)
	input.Split(bufio.ScanWords)
	for input.Scan() {
		word := input.Text()
		wcount[word]++
	}
	if err := input.Err(); err != nil {
		fmt.Fprintf(os.Stderr, "dedup: %v\n", err)
		os.Exit(1)
	}
	fmt.Print("\nword\tcount\n")
	for w, c := range wcount {
		fmt.Printf("%q\t%d\n", w, c)
	}
}
