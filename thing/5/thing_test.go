package thing_test

import (
	"testing"

	"github.com/bitfield/thing"

	"github.com/google/go-cmp/cmp"
)

func TestNewThing_ReturnsThingWithGivenXYZValues(t *testing.T) {
	t.Parallel()
	x, y, z := 1, 2, 3
	want := &thing.Thing{
		X: x,
		Y: y,
		Z: z,
	}
	got, err := thing.NewThing(x, y, z)
	if err != nil {
		t.Fatal(err)
	}
	if !cmp.Equal(want, got) {
		t.Error(cmp.Diff(want, got))
	}
}
