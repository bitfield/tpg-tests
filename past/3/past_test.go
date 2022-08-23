package past3_test

import (
	"testing"
	"time"

	past "past3"

	"github.com/google/go-cmp/cmp"
)

func TestOneHourAgo(t *testing.T) {
	testTime, err := time.Parse(time.RFC3339, "2022-08-05T00:50:25Z")
	if err != nil {
		t.Fatal(err)
	}
	past.Now = func() time.Time {
		return testTime
	}
	want, err := time.Parse(time.RFC3339, "2022-08-04T23:50:25Z")
	if err != nil {
		t.Fatal(err)
	}
	got := past.OneHourAgo()
	if !want.Equal(got) {
		t.Error(cmp.Diff(want, got))
	}
}
