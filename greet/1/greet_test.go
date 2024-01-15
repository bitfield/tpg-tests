package greet_test

import (
	"testing"

	"github.com/bitfield/greet"
)

func TestGreet(t *testing.T) {
	t.Parallel()
	greet.Greet()
}
