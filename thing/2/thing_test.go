package thing2_test

import (
	"testing"
	thing "thing2"
)

func TestNewThing(t *testing.T) {
	t.Parallel()
	got, err := thing.NewThing(1, 2, 3)
	if err != nil {
		t.Fatal(err)
	}
	if got == nil {
		t.Error("want non-nil *Thing, got nil")
	}
}
