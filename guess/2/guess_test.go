//go:build go1.18

package guess2_test

import (
	"testing"

	guess "guess2"
)

func FuzzGuess(f *testing.F) {
	f.Fuzz(func(t *testing.T, input int) {
		guess.Guess(input)
	})
}
