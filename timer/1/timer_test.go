package timer_test

import (
	"testing"
	"time"

	"github.com/bitfield/timer"
)

func TestParseArgs(t *testing.T) {
	t.Parallel()
	want := 10 * time.Millisecond
	got := timer.ParseArgs([]string{"timer", "10ms"})
	if want != got {
		t.Errorf("want %v, got %v", want, got)
	}
}
