package main

import (
	"os"

	hello "hello2"
)

func main() {
	status := hello.Main()
	os.Exit(status)
}
