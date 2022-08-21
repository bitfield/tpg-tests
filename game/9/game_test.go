package game9_test

import (
	"testing"

	game "game9"

	"github.com/google/go-cmp/cmp"
)

func TestListItems(t *testing.T) {
	t.Parallel()
	cases := []struct {
		input []string
		want  string
	}{
		{
			input: []string{},
			want:  "",
		},
		{
			input: []string{
				"a battery",
			},
			want: "You can see a battery here.",
		},
		{
			input: []string{
				"a battery",
				"a key",
			},
			want: "You can see here a battery and a key.",
		},
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
