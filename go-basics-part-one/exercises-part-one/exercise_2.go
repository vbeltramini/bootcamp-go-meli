package exercise

import "fmt"

var temp float64
var moisture uint8
var hPa uint16

func Weather() {
	fmt.Printf("Exercise 2 \n")
	fmt.Printf("\n")

	temp = 28.0
	moisture = 75
	hPa = 1013

	fmt.Printf("Temperatura %v, umidade %v, pressão atmosférica %v hpa", temp, moisture, hPa)
	fmt.Printf("")

}
