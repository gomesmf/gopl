package main

import (
	"fmt"

	"github.com/gomesmf/gopl/ch06-methods/geometry"
)

func main() {
	p := geometry.Point{X: 1, Y: 2}
	q := geometry.Point{X: 4, Y: 6}
	fmt.Println(geometry.Distance(p, q))
	fmt.Println(p.Distance(q))
	fmt.Printf("%T\n", geometry.Distance)
	fmt.Printf("%T\n", p.Distance)

	perim := geometry.Path{
		{X: 1, Y: 1},
		{X: 5, Y: 1},
		{X: 5, Y: 4},
		{X: 1, Y: 1},
	}
	fmt.Println(perim.Distance())
	fmt.Printf("%T %T\n", perim, perim.Distance)

	fmt.Printf("ScaleBy: %T\n", (&p).ScaleBy)

	r := &geometry.Point{X: 1, Y: 2}
	r.ScaleBy(2)
	fmt.Println(*r)
}
