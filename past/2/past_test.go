package past_test

import (
	"testing"
	"time"

	"github.com/bitfield/past"
)

func TestOneHourAgo_ReturnsExpectedTime(t *testing.T) {
	t.Parallel()
	testTime, err := time.Parse(time.RFC3339, "2022-08-05T00:50:25Z")
	if err != nil {
		t.Fatal(err)
	}
	want, err := time.Parse(time.RFC3339, "2022-08-04T23:50:25Z")
	if err != nil {
		t.Fatal(err)
	}
	got := past.OneHourAgo(testTime)
	if !want.Equal(got) {
		t.Errorf("want %v, got %v", want, got)
	}
}
