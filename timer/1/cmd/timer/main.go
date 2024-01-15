package main

import (
	"os"

	"github.com/bitfield/timer"
)

func main() {
	interval := timer.ParseArgs(os.Args)
	timer.Sleep(interval)
}
