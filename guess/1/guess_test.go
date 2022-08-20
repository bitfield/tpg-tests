//go:build go1.18

package guess_test

import (
	"testing"

	"guess"
)

func FuzzGuess(f *testing.F) {
	f.Fuzz(func(t *testing.T, input int) {
		guess.Guess(input)
	})
}
