package ages_test

import (
	"testing"

	"github.com/bitfield/ages"
)

func TestTotalReturnsTotalOfSuppliedAges(t *testing.T) {
	t.Parallel()
	want := 128
	got := ages.Total(makeAgeData())
	if want != got {
		t.Errorf("want %d, got %d", want, got)
	}
}

func makeAgeData() map[string]int {
	return map[string]int{
		"sam":     18,
		"ashley":  72,
		"chandra": 38,
	}
}
