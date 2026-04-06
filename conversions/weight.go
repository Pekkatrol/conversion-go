package conversions

import (
	"fmt"
)

func kilogramsToPounds(kg float64) float64 {
    return kg * 2.20462
}

func poundsToKilograms(pds float64) float64 {
    return pds / 2.20462
}

func Weight() () {
    var value float64
	var choise float64

	fmt.Print("\n1. Kilograms -> Pounds\n")
	fmt.Print("2. Pounds -> Kilograms\n")
	fmt.Print("\n> ")
	fmt.Scan(&choise)
	fmt.Printf("Value : ")
	fmt.Scan(&value)
    switch (choise) {
	    case 1:
            res := kilogramsToPounds(value)
            fmt.Printf("%.2f kg = %.2f lb\n", value, res)
		case 2:
			res := poundsToKilograms(value)
			fmt.Printf("%.2f lb = %.2f kg\n", value, res)
		default:
            fmt.Printf("Invalid choice\n")
	}
}