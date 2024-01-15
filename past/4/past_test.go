package past_test

import (
	"testing"
	"time"

	"github.com/bitfield/past"
)

func TestOneHourAgo(t *testing.T) {
	t.Parallel()
	now := time.Now()
	want := now.Add(-time.Hour)
	got := past.OneHourAgo()
	delta := want.Sub(got).Abs()
	if delta > 10*time.Microsecond {
		t.Errorf("want %v, got %v", want, got)
	}
}
