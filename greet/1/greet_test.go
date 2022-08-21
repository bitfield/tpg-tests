package greet_test

import (
	"testing"

	"greet"
)

func TestGreet(t *testing.T) {
	t.Parallel()
	greet.Greet()
}
