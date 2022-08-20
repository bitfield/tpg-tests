package ages_test

import (
	"testing"

	"ages"

	"github.com/google/go-cmp/cmp"
)

func TestTotalReturnsTotalOfSuppliedAges(t *testing.T) {
	t.Parallel()
	want := 128
	got := ages.Total(makeAgeData())
	if want != got {
		t.Error(cmp.Diff(want, got))
	}
}

func makeAgeData() map[string]int {
	return map[string]int{
		"sam":     18,
		"ashley":  72,
		"chandra": 38,
	}
}
