package main

import "fmt"

// START OMIT
type TemperatureCelsius float64

func (t TemperatureCelsius) ToFahrenheit() float64 {
	return float64(t)*1.8 + 32.0
}

func main() {
	myTemp := TemperatureCelsius(23.7)
	fmt.Printf("°C: %0.2f, °F: %0.2f", myTemp, myTemp.ToFahrenheit())
}

// END OMIT
