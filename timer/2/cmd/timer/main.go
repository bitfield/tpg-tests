package main

import (
	"os"

	timer "timer2"
)

func main() {
	t := timer.NewTimerFromArgs(os.Args)
	t.Sleep()
}
