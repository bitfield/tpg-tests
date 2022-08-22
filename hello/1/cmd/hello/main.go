package main

import (
	"hello"
	"os"
)

func main() {
	status := hello.Main()
	os.Exit(status)
}
