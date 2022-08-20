package thing_test

import (
	"testing"
	"thing"
)

func TestNewThing(t *testing.T) {
	t.Parallel()
	_, err := thing.NewThing(1, 2, 3)
	if err != nil {
		t.Fatal(err)
	}
}
