package main

import (
	"fmt"
	"time"

	"github.com/gomesmf/gopl/ch09-concurrency-shared-variables/bank"
)

func main() {
	// Alice:
	go func() {
		bank.Deposit(200)                // A1
		fmt.Println("=", bank.Balance()) // A2
	}()

	// Bob:
	go bank.Deposit(100) // B
	time.Sleep(10 * time.Millisecond)
}
