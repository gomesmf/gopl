// Ftoc prints two Fahrenheit-to-Celsius conversions.
package main

import "fmt"

const degree string = "\u00b0"

func tempStr(v float64, unit string) string {
	return fmt.Sprintf("%g%s%s", v, degree, unit)
}

func tempF(v float64) string {
	return tempStr(v, "F")
}

func tempC(v float64) string {
	return tempStr(v, "C")
}

func fToC(f float64) float64 {
	return (f - 32) * 5 / 9
}

func main() {
	const freezingF, boilingF = 32.0, 212.0
	fmt.Printf("%s = %s\n", tempF(freezingF), tempC(fToC(freezingF)))
	fmt.Printf("%s = %s\n", tempF(boilingF), tempC(fToC(boilingF)))
}
