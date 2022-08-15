package main

import (
	"fmt"
	"os"
	"time"
)

func launch() {
	fmt.Println("Launch started!")
}

func main() {
	abort := make(chan struct{})
	go func() {
		os.Stdin.Read(make([]byte, 1)) // read a single byte
		abort <- struct{}{}
	}()

	fmt.Println("Commencing countdown. Press return to abort.")

	tick := time.Tick(1 * time.Second)

	for countdown := 5; countdown > 0; countdown-- {
		// fmt.Printf("\r%02d", countdown)
		fmt.Println(countdown)
		select {
		case <-tick:
			// Do nothing
		case <-abort:
			fmt.Println("\nLaunch aborted!")
			return
		}
	}

	launch()
}
