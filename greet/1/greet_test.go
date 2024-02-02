package greet_test

import (
	"testing"

	"github.com/bitfield/greet"
)

func TestGreetAsksForNameAndPrintsHelloName(t *testing.T) {
	t.Parallel()
	greet.Greet()
}
