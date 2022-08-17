// Boiling prints the boiling point of water.
package main

import "fmt"

const degree string = "\u00b0"
const boilingF = 212.0

func tempStr(v float64, unit string) string {
	return fmt.Sprintf("%g%s%s", v, degree, unit)
}

func tempF(v float64) string {
	return tempStr(v, "F")
}

func tempC(v float64) string {
	return tempStr(v, "C")
}

func main() {
	var f = boilingF
	var c = (f - 32) * 5 / 9
	fmt.Printf("boiling point = %s or %s \n", tempF(f), tempC(c))
}
