package thing_test

import (
	"testing"

	"github.com/bitfield/thing"
)

func TestNewThing(t *testing.T) {
	t.Parallel()
	x, y, z := 1, 2, 3
	got, err := thing.NewThing(x, y, z)
	if err != nil {
		t.Fatal(err)
	}
	if got.X != x {
		t.Errorf("want X: %v, got X: %v", x, got.X)
	}
	if got.Y != y {
		t.Errorf("want Y: %v, got Y: %v", y, got.Y)
	}
	if got.Z != z {
		t.Errorf("want Z: %v, got Z: %v", z, got.Z)
	}
}
