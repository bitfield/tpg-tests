//go:build go1.18

package runes2_test

import (
	"testing"
	"unicode/utf8"

	runes "runes2"
)

func FuzzFirstRune(f *testing.F) {
	f.Add("Hello")
	f.Add("world")
	f.Fuzz(func(t *testing.T, s string) {
		got := runes.FirstRune(s)
		want, _ := utf8.DecodeRuneInString(s)
		if want == utf8.RuneError {
			t.Skip() // don't bother testing invalid runes
		}
		if want != got {
			t.Errorf("given %q (0x%[1]x): want '%c' (0x%[2]x)",
				s, want)
			t.Errorf("got '%c' (0x%[1]x)", got)
		}
	})
}
