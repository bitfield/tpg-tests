package timer2_test

import (
	"testing"
	"time"

	timer "timer2"

	"github.com/google/go-cmp/cmp"
)

func TestNewTimerFromArgs(t *testing.T) {
	t.Parallel()
	args := []string{"program", "10s"}
	want := timer.Timer{
		Interval: 10 * time.Second,
	}
	got := timer.NewTimerFromArgs(args)
	if !cmp.Equal(want, got) {
		t.Error(cmp.Diff(want, got))
	}
}
