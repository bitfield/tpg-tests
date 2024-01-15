package game_test

import (
	"testing"

	"github.com/bitfield/game"

	"github.com/google/go-cmp/cmp"
)

func TestListItems_GivesCorrectResultForInput(t *testing.T) {
	t.Parallel()
	type testCase struct {
		input []string
		want  string
	}
	cases := []testCase{
		{
			input: []string{
				"a battery",
				"a key",
				"a tourist map",
			},
			want: "You can see here a battery, a key, and a tourist map.",
		},
	}
	for _, tc := range cases {
		got := game.ListItems(tc.input)
		if tc.want != got {
			t.Error(cmp.Diff(tc.want, got))
		}
	}
}
