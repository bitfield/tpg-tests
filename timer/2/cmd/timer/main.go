package main

import (
	"os"

	"github.com/bitfield/timer"
)

func main() {
	t := timer.NewTimerFromArgs(os.Args)
	t.Sleep()
}
