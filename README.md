# Conversion in Golang

## Goal

This project is a command-line converter written in Go.  
The program shows an interactive menu and lets you convert values between:

- temperature
- distance
- weight

## What I used

- `fmt` for display and user input
- `os` to validate program arguments
- functions for each conversion type
- loops and `switch` statements for menu navigation

## You need Golang

If Go is not installed, run:

```bash
sudo apt update
sudo apt install golang-go
```

## How to run

```bash
go run .
```

> The program expects **no arguments**.

## How it works

When the program starts, it shows a main menu:

1. temperature  
2. distance  
3. weight  
4. exit the program

Then you choose a category, a conversion type, and enter the value to convert.

## Available conversions

### Temperature
- Celsius -> Fahrenheit
- Fahrenheit -> Celsius

### Distance
- Kilometers -> Miles
- Miles -> Kilometers

### Weight
- Kilograms -> Pounds
- Pounds -> Kilograms

## Example

```text
=== Convert ===

Choose a number between 0 and 4

1. temperature
2. distance
3. weight
4. exit the program

> 1

1. Celsius -> Fahrenheit
2. Fahrenheit -> Celsius

> 1
Value : 25
25.00 °C = 77.00 °F
```

## Error handling

The program prints an error message if:

- arguments are provided (none are expected)
- a menu choice is invalid
- a conversion choice is invalid

## Notes

- The app runs in a loop until the user exits.
- After each conversion, the user can continue or stop.

---

Enjoy using the converter!
