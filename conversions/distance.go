package conversions

import (
	"fmt"
)

func kilomettersToMiles(km float64) float64 {
    return km * 0.621371
}

func milesToKilometters(miles float64) float64 {
    return miles / 0.621371
}

func Distance() () {
    var value float64
	var choise float64

	fmt.Print("\n1. Kilometers -> Miles\n")
	fmt.Print("2. Miles -> Kilometers\n")
	fmt.Print("\n> ")
	fmt.Scan(&choise)
	fmt.Printf("Value : ")
	fmt.Scan(&value)
    switch (choise) {
	    case 1:
            res := kilomettersToMiles(value)
            fmt.Printf("%.2f km = %.2f miles\n", value, res)
		case 2:
			res := milesToKilometters(value)
			fmt.Printf("%.2f miles = %.2f km\n", value, res)
		default:
            fmt.Printf("Invalid choice\n")
	}
}