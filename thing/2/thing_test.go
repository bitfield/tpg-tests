package thing_test

import (
	"testing"

	"github.com/bitfield/thing"
)

func TestNewThing_ReturnsNonNilThingAndNoError(t *testing.T) {
	t.Parallel()
	got, err := thing.NewThing(1, 2, 3)
	if err != nil {
		t.Fatal(err)
	}
	if got == nil {
		t.Error("want non-nil *Thing, got nil")
	}
}
