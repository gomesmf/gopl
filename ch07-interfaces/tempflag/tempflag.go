package main

import (
	"flag"
	"fmt"

	"github.com/gomesmf/gopl/ch07-interfaces/tempconv"
)

var temp = tempconv.CelsiusFlag("temp", 20.0, "the temperature")

func main() {
	flag.Parse()
	fmt.Println(*temp)
}
