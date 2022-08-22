package timer_test

import (
	"testing"
	"time"

	"timer"

	"github.com/google/go-cmp/cmp"
)

func TestParseArgs(t *testing.T) {
	t.Parallel()
	want := 10 * time.Millisecond
	got := timer.ParseArgs([]string{"timer", "10ms"})
	if !cmp.Equal(want, got) {
		t.Error(cmp.Diff(want, got))
	}
}
