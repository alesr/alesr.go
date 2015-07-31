package main

import (
	"fmt"
	"os"
	"strings"
)

func main() {
	getTempType()
}

// getTempType asks the user to input the origin type to convert
func getTempType() {
	fmt.Print("\n[1] Celsius\n[2] Fahrenheit\n[3] Kelvin\n\nEnter:")
	var input string
	fmt.Scanln(&input)
	run(input)
}

// run checks the input type and run the correspondent action to it
func run(input string) {
	i := strings.ToLower(input)
	switch {
	case i == "1" || i == "celsius":
		getTemp("celsius")
	case i == "2" || i == "fahrenheit":
		getTemp("fahrenheit")
	case i == "3" || i == "kelvin":
		getTemp("kelvin")
	default:
		fmt.Println("Error. Type 1, 2 or 3 and press enter.")
		getTempType()
	}
}

// getTemp asks for the temperature value and call the appropriate conversor
func getTemp(unit string) {
	switch unit {
	case "celsius":
		fmt.Print("\nType the temperature in Celsius:")
		temp, cf, ck := celsiusTo(getInputValue())
		printConversion(unit, temp, cf, ck)
	case "fahrenheit":
		fmt.Print("\nType the temperature in Fahrenheit:")
		temp, fc, fk := fahrenheitTo(getInputValue())
		printConversion(unit, temp, fc, fk)
	case "kelvin":
		fmt.Print("\nType the temperature in Kelvin:")
		temp, kc, kf := kelvinTo(getInputValue())
		printConversion(unit, temp, kc, kf)
	}
}

// getInputValue gets the user input for temperature as a float64 valid value
func getInputValue() float64 {
	var temp float64
	_, err := fmt.Scanf("%f", &temp)
	if err != nil {
		fmt.Println("\nYou must time a valid temperature value.")
		os.Exit(1)
	}
	return temp
}

// celsiusTo fahrenheit and kelvin
func celsiusTo(temp float64) (float64, float64, float64) {
	cf := (temp * 9 / 5) + 32
	ck := temp + 273.15
	return temp, cf, ck
}

// fahrenheitTo celsius and kelvin
func fahrenheitTo(temp float64) (float64, float64, float64) {
	fc := (temp - 32) * 5 / 9
	fk := (temp + 459.67) * 5 / 9
	return temp, fc, fk
}

// kelvinTo celsius and fahrenheit
func kelvinTo(temp float64) (float64, float64, float64) {
	kc := temp - 273.15
	kf := (temp-273.15)*1.8000 + 32.00
	return temp, kc, kf
}

// printConversion based on the type
func printConversion(unit string, temp, a, b float64) {
	switch unit {
	case "celsius":
		fmt.Printf("\nCelsius: %2.f\nFahrenheit: %2.f\nKelvin: %2.f\n", temp, a, b)
	case "fahrenheit":
		fmt.Printf("\nFahrenheit: %2.f\nCelsius: %2.f\nKelvin: %2.f\n", temp, a, b)
	case "kelvin":
		fmt.Printf("\nKelvin: %2.f\nCelsius: %2.f\nFahrenheit: %2.f\n", temp, a, b)
	}
}
