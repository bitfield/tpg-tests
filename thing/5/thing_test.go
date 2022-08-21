package thing5_test

import (
	"testing"
	thing "thing5"

	"github.com/google/go-cmp/cmp"
)

func TestNewThing(t *testing.T) {
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
