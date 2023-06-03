package main

import (
	"fmt"
	"os"
	"strconv"
)

// converte de fahrenheit para celcius
func convToFahrenheit(value float64) {
	fmt.Println("Value in Fahrenheit = ", 1.8*value+32)
}

// converte de celcius para fahrenheit
func convToCelcius(value float64) {
	fmt.Println("Value in Celcius = ", (value-32)/1.8)
}

// convete de milhas para km
func convToMiles(value float64) {
	fmt.Println("Value in miles = ", value/1.609)
}

// converte de km para milhas
func convToKm(value float64) {
	fmt.Println("Value in km = ", value*1.609)
}

// trata linha de comandos
func argparse() {

	value, errorConv := strconv.ParseFloat(os.Args[2], 64)

	if os.Args[1] == "celcius" {
		convToFahrenheit(value)
	} else if os.Args[1] == "fahrenheit" {
		convToCelcius(value)
	} else if os.Args[1] == "km" {
		convToMiles(value)
	} else if os.Args[1] == "miles" {
		convToKm(value)
	} else {
		if errorConv != nil {
			fmt.Println(errorConv)
		}
		fmt.Println("Usage: go run conversor.go km 23")
		os.Exit(-1)
	}
}

func main() {
	argparse()
}
