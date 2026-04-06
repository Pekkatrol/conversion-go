package conversions

import (
	"fmt"
)

func fahrenheitToCelsius(f float64) (float64) {
	return (f-32) * 5/9
}

func celsiusToFahrenheit(c float64) float64 {
    return c * 9 / 5 + 32
}

func Temperature() () {
    var value float64
	var choise float64

	fmt.Print("\n1. Celsius -> Fahrenheit\n")
	fmt.Print("2. Fahrenheit -> Celsius\n")
	fmt.Print("\n> ")
	fmt.Scan(&choise)
	fmt.Printf("Value : ")
	fmt.Scan(&value)
	switch (choise) {
	    case 1:
            res := celsiusToFahrenheit(value)
            fmt.Printf("%.2f °C = %.2f °F\n", value, res)
		case 2:
			res := fahrenheitToCelsius(value)
			fmt.Printf("%.2f °F = %.2f °C\n", value, res)
		default:
            fmt.Printf("Invalid choice\n")
	}
}