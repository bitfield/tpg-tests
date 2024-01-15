package valid_test

import (
	"testing"

	"github.com/bitfield/valid"
)

func TestValidIsTrueForValidInput(t *testing.T) {
	t.Parallel()
	if !valid.Valid("valid input") {
		t.Error(false)
	}
}
