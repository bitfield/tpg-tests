package game_test

import (
	"testing"

	"github.com/bitfield/game"

	"github.com/google/go-cmp/cmp"
)

func TestListItems_GivesCorrectResultForInput(t *testing.T) {
	t.Parallel()
	input := []string{
		"a battery",
		"a key",
		"a tourist map",
	}
	want := "You can see here a battery, a key, and a tourist map."
	got := game.ListItems(input)
	if want != got {
		t.Error(cmp.Diff(want, got))
	}
}
