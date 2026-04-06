package main

import (
	"os"
)

func main() {
	args := os.Args
	if (len(args) != 1) {
        print_error("No arguments needed\n")
	}
	conversion_loop()
}