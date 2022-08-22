package timer2

import (
	"fmt"
	"time"
)

type Timer struct {
	Interval time.Duration
}

func (t Timer) Sleep() {
	fmt.Printf("Starting a timer for %s...\n", t.Interval)
	time.Sleep(t.Interval)
	fmt.Println("...done")
}

func NewTimerFromArgs(args []string) Timer {
	if len(args) < 2 {
		return Timer{}
	}
	interval, err := time.ParseDuration(args[1])
	if err != nil {
		return Timer{}
	}
	return Timer{
		Interval: interval,
	}
}
