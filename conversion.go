package main

import (
	"fmt"
	"conversion/conversions"
)

const (
    TEMPERATURE = iota + 1
    DISTANCE
    WEIGHT
    EXIT
	LEN
)

func print_choises() {
	fmt.Printf("\033[2J\033[H")
    fmt.Printf("=== Convert ===\n\n")
    fmt.Printf("Choose a number between 0 and %d\n\n", LEN - 1)
    fmt.Printf("%d. temperature\n", TEMPERATURE)
    fmt.Printf("%d. distance\n", DISTANCE)
    fmt.Printf("%d. weight\n", WEIGHT)
    fmt.Printf("%d. exit the program\n", EXIT)
    fmt.Print("\n> ")
}

func conversion(value int) () {
	switch (value) {
		case TEMPERATURE:
			conversions.Temperature()
		case DISTANCE:
			conversions.Distance()
		case WEIGHT:
			conversions.Weight()
		case EXIT:
			fmt.Printf("Bye\n")
		default:
			fmt.Printf("Wrong choice\n")
	}
}

func conversion_loop() () {
	var value int
	var again int
	
	for {
		print_choises()
		fmt.Scan(&value)
        conversion(value)
		if (value == EXIT) {
			break
		}
		fmt.Printf("\nDo you want to continue ?\n")
		fmt.Printf("0. No\n1. Yes\n")
		fmt.Scan(&again)
		if (again == 0) {
			break
		}
	}
}