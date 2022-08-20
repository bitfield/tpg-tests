package square_test

import (
	"math/rand"
	"strconv"
	"testing"

	"square"
)

func TestSquareResultIsAlwaysNonNegative(t *testing.T) {
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
