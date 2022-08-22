package main

import (
	"os"

	greet "greet7"
)

func main() {
	status := greet.Main()
	os.Exit(status)
}
