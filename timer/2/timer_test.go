package timer_test

import (
	"testing"
	"time"

	"github.com/bitfield/timer"
)

func TestNewTimerFromArgs(t *testing.T) {
	t.Parallel()
	args := []string{"program", "10s"}
	want := timer.Timer{
		Interval: 10 * time.Second,
	}
	got := timer.NewTimerFromArgs(args)
	if want != got {
		t.Errorf("want %v, got %v", want, got)
	}
}
