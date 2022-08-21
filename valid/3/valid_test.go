package valid3_test

import (
	"testing"
	valid "valid3"
)

func TestValidIsTrueForValidInput(t *testing.T) {
	t.Parallel()
	if !valid.Valid("valid input") {
		t.Error(false)
	}
}
