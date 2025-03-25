package square_test

import (
	"math/rand/v2"
	"strconv"
	"testing"

	"github.com/bitfield/square"
)

func TestSquareGivesNonNegativeResult(t *testing.T) {
	t.Parallel()
	inputs := rand.Perm(100)
	for _, n := range inputs {
		t.Run(strconv.Itoa(n), func(t *testing.T) {
			got := square.Square(n)
			if got < 0 {
				t.Errorf("Square(%d) is negative: %d", n, got)
			}
		})
	}
}
