package guess_test

import (
	"testing"

	"github.com/bitfield/guess"
)

func FuzzGuess(f *testing.F) {
	f.Fuzz(func(t *testing.T, input int) {
		guess.Guess(input)
	})
}
