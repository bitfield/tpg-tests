package main

import (
	"os"

	"timer"
)

func main() {
	interval := timer.ParseArgs(os.Args)
	timer.Sleep(interval)
}
