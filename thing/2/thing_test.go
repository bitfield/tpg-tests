package thing_test

import (
	"testing"
	"thing"
)

func TestNewThing(t *testing.T) {
	got, err := thing.NewThing(1, 2, 3)
	if err != nil {
		t.Fatal(err)
	}
	if got == nil {
		t.Error("want non-nil *Thing, got nil")
	}
}
