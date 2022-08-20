package game_test

import (
	"testing"

	"game"

	"github.com/google/go-cmp/cmp"
)

func TestListItems(t *testing.T) {
	t.Parallel()
	cases := map[string]struct {
		input []string
		want  string
	}{
		"no items": {
			input: []string{},
			want:  "",
		},
		"one item": {
			input: []string{
				"a battery",
			},
			want: "You can see a battery here.",
		},
		"two items": {
			input: []string{
				"a battery",
				"a key",
			},
			want: "You can see here a battery and a key.",
		},
		"three items": {
			input: []string{
				"a battery",
				"a key",
				"a tourist map",
			},
			want: "You can see here a battery, a key, and a tourist map.",
		},
	}
	for name, tc := range cases {
		t.Run(name, func(t *testing.T) {
			got := game.ListItems(tc.input)
			if tc.want != got {
				t.Error(cmp.Diff(tc.want, got))
			}
		})
	}
}
