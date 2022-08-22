package main

import (
	"os"

	greet "greet6"
)

func main() {
	status := greet.Main()
	os.Exit(status)
}
