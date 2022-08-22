package timer

import (
	"fmt"
	"time"
)

func ParseArgs(args []string) time.Duration {
	if len(args) < 2 {
		return 0
	}
	interval, err := time.ParseDuration(args[1])
	if err != nil {
		return 0
	}
	return interval
}

func Sleep(interval time.Duration) {
	fmt.Printf("Starting a timer for %s...\n", interval)
	time.Sleep(interval)
	fmt.Println("...done")
}
